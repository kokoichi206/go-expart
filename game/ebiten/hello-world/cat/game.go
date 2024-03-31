package cat

import (
	"context"
	"fmt"
	"image/color"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"time"

	"hello-world/cat/touch"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Game struct {
	closer func()
	logger *slog.Logger

	cat    *Cat
	snakes []*Snake
	tm     *touch.TouchManager

	enemyCh chan *enemy

	score     int
	highScore int
}

func NewGame(logLevel string) *Game {
	lv := slog.LevelInfo
	switch logLevel {
	case "debug", "DEBUG":
		lv = slog.LevelDebug
	default:
	}
	h := slog.NewJSONHandler(
		os.Stderr,
		&slog.HandlerOptions{
			Level: lv,
		},
	)
	slog.SetDefault(slog.New(h))
	logger := slog.Default()

	go goroutineCheck(logger)

	g := &Game{
		logger:  logger,
		enemyCh: make(chan *enemy),
		tm:      touch.NewTouchManager(),
	}
	g.initGame()
	return g
}

func (g *Game) initGame() {
	if g.closer != nil {
		g.closer()
	}

	g.cat = NewCat(g.tm)
	g.snakes = []*Snake{}
	ctx, cancel := context.WithCancel(context.Background())
	g.closer = func() {
		cancel()
	}
	g.score = 0
	g.highScore = max(g.highScore, g.score)

	go g.enemyGenerator(ctx)
}

func (g *Game) enemyGenerator(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second * time.Duration(2+rand.Intn(3))):
			// 2秒から4秒のランダムな遅延を生成。
		}

		g.enemyCh <- &enemy{
			enType: snake,
		}
	}
}

func (g *Game) Update() error {
	g.tm.Update()
	if g.dead() {
		if ebiten.IsKeyPressed(ebiten.KeySpace) ||
			(touch.IsTouchMain() && g.tm.IsJustTouched()) {
			g.initGame()
		}
		return nil
	}

	g.cat.update()
	for _, s := range g.snakes {
		s.update()
	}

	g.updateStage()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("x, y, vx, vy: %.2f, %.2f, %.2f, %.2f", g.cat.pos.X, g.cat.pos.Y, g.cat.vec.X, g.cat.vec.Y))

	g.showScore(screen)

	g.cat.draw(screen)
	for _, s := range g.snakes {
		s.draw(screen)
	}

	if g.dead() {
		showGameOver(screen)
	}
}

func (g *Game) showScore(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("%2d", g.score), mpFont, ScreenWidth-50, 30, color.White)
}

func (g *Game) dead() bool {
	for _, s := range g.snakes {
		// 球体とみなして当たり判定を行う。
		if math.Pow(g.cat.pos.X-s.pos.X, 2)+math.Pow(g.cat.pos.Y-s.pos.Y, 2) < math.Pow(catSize/2+enemySize/2, 2) {
			return true
		}
	}
	return false
}

func (g *Game) updateStage() {
	select {
	case e := <-g.enemyCh:
		switch e.enType {
		case snake:
			g.logger.Debug("new snake created")
			g.snakes = append(g.snakes, NewSnake())
		default:
		}
	default:
	}

	for _, s := range g.snakes {
		if s.pos.X < -enemyImgSize {
			g.score++
			g.logger.Debug(fmt.Sprintf("score up! (new score: %d)", g.score))
			g.snakes = g.snakes[1:]
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func showGameOver(screen *ebiten.Image) {
	var restartStr string
	if touch.IsTouchMain() {
		restartStr = "Touch to restart"
	} else {
		restartStr = "Press 'Space' to restart"
	}

	b, _ := font.BoundString(mpFont, restartStr)
	text.Draw(screen, restartStr, mpFont, ScreenWidth/2-b.Max.X.Round()/2, 80, color.RGBA{0xff, 0x00, 0x00, 0xFF})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ScreenWidth/2-float64(gameoverLogo.Bounds().Dx())/2, ScreenHeight/2-float64(gameoverLogo.Bounds().Dy())/2)
	screen.DrawImage(gameoverLogo, op)
}
