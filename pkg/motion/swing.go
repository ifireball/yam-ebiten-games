package motion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Swing struct {
	Pivot      gmath.Vec2
	MinAngle   float32
	MaxAngle   float32
	StartAngle float32
	Duration   float32
	Cycles     int
}

func (s *Swing) Run() StepFunc {
	seq := gween.NewSequence(
		gween.New(s.StartAngle, s.MinAngle, s.Duration/2, ease.InOutQuart),
		gween.New(s.MinAngle, s.MaxAngle, s.Duration, ease.InOutQuart),
		gween.New(s.MaxAngle, s.MinAngle, s.Duration, ease.InOutQuart),
		gween.New(s.MinAngle, s.StartAngle, s.Duration/2, ease.InOutQuart),
	)

	return func(step *ebiten.GeoM) bool {
		angle, _, done := seq.Update(1)
		step.Reset()
		if done {
			return false
		}
		step.Translate(-s.Pivot.X, -s.Pivot.Y)
		step.Rotate(float64(angle))
		step.Translate(s.Pivot.Unwrap())
		return true
	}
}
