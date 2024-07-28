package girl

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
)

type ArrowKeysCtrl struct{}

func (*ArrowKeysCtrl) Control(position *float64) {
	switch {
	case keyboard.IsPressedOneOf(keyboard.WobeeYellow, ebiten.KeyLeft):
		*position = *position - Speed
	case keyboard.IsPressedOneOf(keyboard.WobeeRed, ebiten.KeyRight):
		*position = *position + Speed
	}
}
