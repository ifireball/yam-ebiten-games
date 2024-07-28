package girl

import (
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/protocols"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type AutoCtrl struct {
	Fruit     protocols.WithActiveRect
	direction float64
}

func (ac *AutoCtrl) Control(position *float64) {
	var ar gmath.Rect
	ac.Fruit.GetActiveRect(&ar)
	if !ar.IsZero() && ar.Top() < ScreenBasketBottom {
		activeCenter := (ar.TopLeft.X + ar.BottomRight.X) / 2
		switch {
		case *position > activeCenter+BasketWidth/2:
			ac.direction = -1
		case *position < activeCenter-BasketWidth/2:
			ac.direction = 1
		}

	}

	*position += ac.direction * Speed
	if *position < MinPosition {
		ac.direction = 1
	}
	if *position > MaxPosition {
		ac.direction = -1
	}

}
