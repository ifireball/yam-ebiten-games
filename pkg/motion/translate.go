package motion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Translate struct {
	From, To gmath.Vec2
	Duration float32
	Easing   ease.TweenFunc
}

func (t *Translate) Run() StepFunc {
	xt := gween.New(float32(t.From.X), float32(t.To.X), t.Duration, t.Easing)
	yt := gween.New(float32(t.From.Y), float32(t.To.Y), t.Duration, t.Easing)

	return func(step *ebiten.GeoM) bool {
		x, doneX := xt.Update(1)
		y, doneY := yt.Update(1)
		if doneX || doneY {
			return false
		}
		step.Reset()
		step.Translate(float64(x), float64(y))
		return true
	}
}
