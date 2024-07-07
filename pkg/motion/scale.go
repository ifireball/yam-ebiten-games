package motion

import (
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Scale struct {
	From, To, Pivot gmath.Vec2
	Duration float32
	Easing ease.TweenFunc
}

func (s *Scale) Run(out chan<- Step) {
	defer lastStep(out)

	xt := gween.New(float32(s.From.X), float32(s.To.X), s.Duration, s.Easing)
	yt := gween.New(float32(s.From.Y), float32(s.To.Y), s.Duration, s.Easing)
	var step Step
	for {
		x, doneX := xt.Update(1)
		y, doneY := yt.Update(1)
		if doneX || doneY {
			break
		}
		step.Transform.Reset()
		step.Transform.Translate(-s.Pivot.X, -s.Pivot.Y)
		step.Transform.Scale(float64(x), float64(y))
		step.Transform.Translate(s.Pivot.Unwrap())
		out <- step
	}
}
