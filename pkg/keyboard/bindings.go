package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WobeeBlue = ebiten.KeyS
	WobeeYellow = ebiten.KeyG
	WobeeRed = ebiten.KeyK
	WobeeGreen = ebiten.KeyApostrophe
)

var WobeeBlueGreen = NewCombo(WobeeBlue, WobeeGreen)
