package girl

import (
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/protocols"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type AutoCtrl struct {
	Fruit protocols.WithActiveRect
}

func (ac *AutoCtrl) Control(position *float64) {
	var ar gmath.Rect
	ac.Fruit.GetActiveRect(&ar)
	if ar.IsZero() {
		return
	}

	activeCenter := (ar.TopLeft.X + ar.BottomRight.X) / 2
	switch {
	case *position > activeCenter + BasketWidth/2:
		*position -= Speed
	case *position < activeCenter - BasketWidth/2:
		*position += Speed
	}
}
