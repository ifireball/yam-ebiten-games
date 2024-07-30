package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WobeeBlue   = ebiten.KeyS
	WobeeYellow = ebiten.KeyG
	WobeeRed    = ebiten.KeyK
	WobeeGreen  = ebiten.KeyApostrophe
)

const WobeeKeysAmount = 4

var WobeeOrder = [WobeeKeysAmount]ebiten.Key{WobeeBlue, WobeeYellow, WobeeRed, WobeeGreen}
var WobeeColorNames = [WobeeKeysAmount]string{"blue", "yellow", "red", "green"}

var WobeeBlueGreen = NewCombo(WobeeBlue, WobeeGreen)
