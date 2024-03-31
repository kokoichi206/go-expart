package touch

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TouchManager struct {
	justTouchedIDs map[ebiten.TouchID]struct{}
}

func NewTouchManager() *TouchManager {
	return &TouchManager{
		justTouchedIDs: map[ebiten.TouchID]struct{}{},
	}
}

func (tm *TouchManager) Update() {
	tm.justTouchedIDs = map[ebiten.TouchID]struct{}{}

	// Touch
	for _, touch := range inpututil.AppendJustPressedTouchIDs(nil) {
		// You can get the touch position by ebiten.TouchPosition if you need.
		// x, y := ebiten.TouchPosition(touch)
		tm.justTouchedIDs[touch] = struct{}{}
	}
}

func (tm *TouchManager) IsJustTouched() bool {
	return len(tm.justTouchedIDs) > 0
}

func IsTouchMain() bool {
	return isTouchMain()
}
