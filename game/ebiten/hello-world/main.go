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
)

var (
	//go:embed cat.png
	Cat_png []byte

	catImage *ebiten.Image
	bg       *ebiten.Image
)

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
		},
	}
}

type Cat struct {
	// 中心のポジション。
	pos Position
}

func (c *Cat) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X-catSize/2), float64(screenHeight-(c.pos.Y+catSize/2)))
	screen.DrawImage(catImage, op)
}

type Position struct {
	X int
	Y int
}

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

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("x, y: %d, %d", g.cat.pos.X, g.cat.pos.Y))

	screen.DrawImage(bg, nil)

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
