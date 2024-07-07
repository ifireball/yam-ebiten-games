package motion

import (
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Trnaslate struct {
	From, To gmath.Vec2
	Duration float32
	Easing ease.TweenFunc
}

func (t *Trnaslate) Run(out chan<- Step) {
	defer lastStep(out)

	xt := gween.New(float32(t.From.X), float32(t.To.X), t.Duration, t.Easing)
	yt := gween.New(float32(t.From.Y), float32(t.To.Y), t.Duration, t.Easing)
	var step Step
	for {
		x, doneX := xt.Update(1)
		y, doneY := yt.Update(1)
		if doneX || doneY {
			break
		}
		step.Transform.Reset()
		step.Transform.Translate(float64(x), float64(y))
		out <- step
	}
}
