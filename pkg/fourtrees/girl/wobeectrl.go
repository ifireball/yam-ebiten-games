package girl

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/protocols"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type WobeeCtrl struct {
	Fruit  protocols.WithActiveType
	posIdx int
	target float64
	tween  *gween.Tween
}

var keyToPos = map[ebiten.Key]int{
	keyboard.WobeeBlue:   0,
	keyboard.WobeeYellow: 1,
	keyboard.WobeeRed:    2,
	keyboard.WobeeGreen:  3,
}

func (wc *WobeeCtrl) Control(position *float64) {
	var ar gmath.Rect
	var target float64

	wc.Fruit.GetActiveRect(&ar)

	if newPos, ok := positionJustPressed(); ok {
		wc.posIdx = newPos
	}
	if wc.posIdx == wc.Fruit.GetActiveType() && ar.Top() < ScreenBasketBottom {
		target = (ar.Left() + ar.Right()) / 2
	} else {
		target = gdata.ScreenWidth * float64(wc.posIdx+1) / 5
	}
	if wc.target != target {
		wc.target = target
		wc.tween = gween.New(float32(*position), float32(wc.target), 20, ease.InOutCubic)
	}
	if wc.tween != nil {
		pos32, isFinished := wc.tween.Update(1)
		*position = float64(pos32)
		if isFinished {
			wc.tween = nil
		}
	}
}

func positionJustPressed() (int, bool) {
	for key, posIdx := range keyToPos {
		if inpututil.IsKeyJustPressed(key) {
			return posIdx, true
		}
	}
	return -1, false
}
