package mobile

import (
	"hello-world/cat"

	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	g := cat.NewGame("debug")
	mobile.SetGame(g)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
