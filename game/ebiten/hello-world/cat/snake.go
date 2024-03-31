package cat

import "github.com/hajimehoshi/ebiten/v2"

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
	op.GeoM.Translate(float64(s.pos.X-enemyImgSize/2), float64(ScreenHeight-(s.pos.Y+enemyImgSize/2)))
	screen.DrawImage(enemySnakeImage, op)
}

func NewSnake() *Snake {
	return &Snake{
		pos: Position{
			X: ScreenWidth,
			Y: enemyImgSize / 2,
		},
		vec: Velocity{
			X: 0,
			Y: 0,
		},
	}
}
