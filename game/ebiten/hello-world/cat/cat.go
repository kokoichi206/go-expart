package cat

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"hello-world/cat/touch"
)

const (
	maxJumpStep = 3
)

type Cat struct {
	// 中心のポジション。
	pos Position
	vec Velocity

	jumpStep int

	tm *touch.TouchManager
}

func NewCat(tm *touch.TouchManager) *Cat {
	return &Cat{
		pos: Position{
			X: ScreenWidth / 3,
			Y: catImgSize / 2,
		},
		vec: Velocity{
			X: 0,
			Y: 0,
		},
		tm: tm,
	}
}

func (c *Cat) update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.pos.X -= catSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.pos.X += catSpeed
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if c.jumpStep < maxJumpStep {
			c.vec.Y += 5
			c.jumpStep += 1	
		}
	}

	if touch.IsTouchMain() && c.tm.IsJustTouched() {
		if c.jumpStep < maxJumpStep {
			c.vec.Y += 5
			c.jumpStep += 1	
		}
	}

	// y 軸のバリデーション。
	if c.pos.Y+c.vec.Y <= catImgSize/2 {
		c.vec.Y = 0
		c.pos.Y = catImgSize / 2
		// 着地時に step カウントをリセット。
		c.jumpStep = 0
	} else {
		c.pos.Y += c.vec.Y
		c.vec.Y -= 0.1
	}

	// x 軸のバリデーション。
	if c.pos.X < catImgSize/2 {
		c.pos.X = catImgSize / 2
	}
	if c.pos.X > ScreenWidth-catImgSize/2 {
		c.pos.X = ScreenWidth - catImgSize/2
	}
}

func (c *Cat) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.pos.X-catImgSize/2), float64(ScreenHeight-(c.pos.Y+catImgSize/2)))
	screen.DrawImage(catImage, op)
}
