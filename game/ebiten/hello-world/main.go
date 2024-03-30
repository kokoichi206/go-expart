package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 640
	screenHeight = 480

	catImgSize = 64
	catSize    = 48 // 当たり判定に使う、実際の猫のサイズ。
	catSpeed   = 4

	enemyImgSize = 64
	enemySize    = 48 // 当たり判定に使う、実際の敵のサイズ。
	snakeSpeed   = 1.5
)

var (
	//go:embed cat.png
	Cat_png []byte

	//go:embed snake.png
	Snake_png []byte

	//go:embed gameover.png
	Gameover_png []byte

	catImage        *ebiten.Image
	enemySnakeImage *ebiten.Image

	bg           *ebiten.Image
	gameoverLogo *ebiten.Image

	mpFont font.Face
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(Cat_png))
	if err != nil {
		log.Fatal("image.Decode error: ", err)
	}
	catImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(Snake_png))
	if err != nil {
		log.Fatal("image.Decode error: ", err)
	}
	enemySnakeImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(Gameover_png))
	if err != nil {
		log.Fatal("image.Decode error: ", err)
	}
	gameoverLogo = ebiten.NewImageFromImage(img)

	bg = ebiten.NewImage(screenWidth, screenHeight)
	vector.DrawFilledRect(
		bg, 0, 0, screenWidth, screenHeight,
		color.RGBA{0xaa, 0xaa, 0xaa, 0xff}, true)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mpFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	closer func()
	logger *slog.Logger

	cat    *Cat
	snakes []*Snake

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
		logger: logger,
		// enemyCh: make(chan *enemy, 3),
		enemyCh: make(chan *enemy),
	}
	g.initGame()
	return g
}

func goroutineCheck(logger *slog.Logger) {
	ticker := time.NewTicker(4 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		logger.Debug("goroutine check",
			slog.Attr{
				Key:   "count",
				Value: slog.IntValue(runtime.NumGoroutine()),
			},
		)
	}
}

func (g *Game) initGame() {
	if g.closer != nil {
		g.closer()
	}

	g.cat = NewCat()
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

type Cat struct {
	// 中心のポジション。
	pos Position
	vec Velocity
}

func NewCat() *Cat {
	return &Cat{
		pos: Position{
			X: screenWidth / 3,
			Y: catImgSize / 2,
		},
		vec: Velocity{
			X: 0,
			Y: 0,
		},
	}
}

func (c *Cat) update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.pos.X -= catSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.pos.X += catSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeySpace)) &&
		c.vec.Y == 0 {
		c.vec.Y += 4
	}

	// y 軸のバリデーション。
	if c.pos.Y+c.vec.Y <= catImgSize/2 {
		c.vec.Y = 0
		c.pos.Y = catImgSize / 2
	} else {
		c.pos.Y += c.vec.Y
		c.vec.Y -= 0.1
	}

	// x 軸のバリデーション。
	if c.pos.X < catImgSize/2 {
		c.pos.X = catImgSize / 2
	}
	if c.pos.X > screenWidth-catImgSize/2 {
		c.pos.X = screenWidth - catImgSize/2
	}
}

func (c *Cat) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X-catImgSize/2), float64(screenHeight-(c.pos.Y+catImgSize/2)))
	screen.DrawImage(catImage, op)
}

type enemy struct {
	enType enType
}

type enType int

const (
	snake enType = iota
)

type Snake struct {
	// 中心のポジション。
	pos Position
	vec Velocity
}

func (s *Snake) update() {
	s.pos.X -= snakeSpeed
}

func (s *Snake) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.pos.X-enemyImgSize/2), float64(screenHeight-(s.pos.Y+enemyImgSize/2)))
	screen.DrawImage(enemySnakeImage, op)
}

func NewSnake() *Snake {
	return &Snake{
		pos: Position{
			X: screenWidth,
			Y: enemyImgSize / 2,
		},
		vec: Velocity{
			X: 0,
			Y: 0,
		},
	}
}

type Position struct {
	X float64
	Y float64
}

type Velocity struct {
	X float64
	Y float64
}

func (g *Game) Update() error {
	if g.dead() {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
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
	text.Draw(screen, fmt.Sprintf("%2d", g.score), mpFont, screenWidth-50, 30, color.White)
}

func showGameOver(screen *ebiten.Image) {
	const restartStr = "Press 'Space' to restart"
	b, _ := font.BoundString(mpFont, restartStr)
	text.Draw(screen, restartStr, mpFont, screenWidth/2-b.Max.X.Round()/2, 80, color.RGBA{0xff, 0x00, 0x00, 0xFF})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(screenWidth/2-float64(gameoverLogo.Bounds().Dx())/2, screenHeight/2-float64(gameoverLogo.Bounds().Dy())/2)
	screen.DrawImage(gameoverLogo, op)
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
	return screenWidth, screenHeight
}

func main() {
	lv := os.Getenv("LOG_LEVEL")

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cat game")
	if err := ebiten.RunGame(NewGame(lv)); err != nil {
		log.Fatal(err)
	}
}
