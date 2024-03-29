package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	catSize      = 64
	catSpeed     = 4
)

var (
	//go:embed cat.png
	Cat_png []byte

	catImage *ebiten.Image
	bg       *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(Cat_png))
	if err != nil {
		log.Fatal("image.Decode error: ", err)
	}

	bg = ebiten.NewImage(screenWidth, screenHeight)
	vector.DrawFilledRect(
		bg, 0, 0, screenWidth, screenHeight,
		color.RGBA{0, 0xff, 0, 0xff}, true)

	catImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	cat *Cat
}

func NewGame() *Game {
	return &Game{
		cat: &Cat{
			pos: Position{
				X: screenWidth / 3,
				Y: catSize / 2,
			},
			vec: Velocity{
				X: 0,
				Y: 0,
			},
		},
	}
}

type Cat struct {
	// 中心のポジション。
	pos Position
	vec Velocity
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
	if c.pos.Y+c.vec.Y <= catSize/2 {
		c.vec.Y = 0
		c.pos.Y = catSize / 2
	} else {
		c.pos.Y += c.vec.Y
		c.vec.Y -= 0.1
	}
}

func (c *Cat) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X-catSize/2), float64(screenHeight-(c.pos.Y+catSize/2)))
	screen.DrawImage(catImage, op)
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
	g.cat.update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bg, nil)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("x, y, vx, vy: %.2f, %.2f, %.2f, %.2f", g.cat.pos.X, g.cat.pos.Y, g.cat.vec.X, g.cat.vec.Y))
	g.cat.draw(screen)
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
