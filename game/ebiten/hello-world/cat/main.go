package cat

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"log/slog"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480

	catImgSize = 64
	catSize    = 48 // 当たり判定に使う、実際の猫のサイズ。
	catSpeed   = 4

	enemyImgSize = 64
	enemySize    = 48 // 当たり判定に使う、実際の敵のサイズ。
	snakeSpeed   = 1.5
)

var (
	//go:embed assets/cat.png
	Cat_png []byte

	//go:embed assets/snake.png
	Snake_png []byte

	//go:embed assets/gameover.png
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

	bg = ebiten.NewImage(ScreenWidth, ScreenHeight)
	vector.DrawFilledRect(
		bg, 0, 0, ScreenWidth, ScreenHeight,
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

type Position struct {
	X float64
	Y float64
}

type Velocity struct {
	X float64
	Y float64
}
