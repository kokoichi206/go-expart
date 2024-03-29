package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		color.RGBA{0, 0xff, 0, 0xff}, true)
}

type Game struct {
	cat   *Cat
	snake *Snake
}

func NewGame() *Game {
	return &Game{
		cat:   NewCat(),
		snake: NewSnake(),
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

	// 地面より下に行かないようにする。
	if c.pos.Y+c.vec.Y <= catImgSize/2 {
		c.vec.Y = 0
		c.pos.Y = catImgSize / 2
	} else {
		c.pos.Y += c.vec.Y
		c.vec.Y -= 0.1
	}
}

func (c *Cat) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X-catImgSize/2), float64(screenHeight-(c.pos.Y+catImgSize/2)))
	screen.DrawImage(catImage, op)
}

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
	if !g.dead() {
		g.cat.update()
		g.snake.update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("x, y, vx, vy: %.2f, %.2f, %.2f, %.2f", g.cat.pos.X, g.cat.pos.Y, g.cat.vec.X, g.cat.vec.Y))

	g.cat.draw(screen)
	g.snake.draw(screen)

	if g.dead() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(screenWidth/2-float64(gameoverLogo.Bounds().Dx())/2, screenHeight/2-float64(gameoverLogo.Bounds().Dy())/2)
		screen.DrawImage(gameoverLogo, op)
	}
}

func (g *Game) dead() bool {
	// 球体とみなして当たり判定を行う。
	return math.Pow(g.cat.pos.X-g.snake.pos.X, 2)+math.Pow(g.cat.pos.Y-g.snake.pos.Y, 2) < math.Pow(catSize/2+enemySize/2, 2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Cat game")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
