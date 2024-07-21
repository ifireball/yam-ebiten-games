package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WobeeBlue = ebiten.KeyA
	WobeeYellow = ebiten.KeyF
	WobeeRed = ebiten.KeyJ
	WobeeGreen = ebiten.KeySemicolon
)

var WobeeBlueGreen = NewCombo(WobeeBlue, WobeeGreen)
