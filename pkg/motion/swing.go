package motion

import (
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type Swing struct {
	Pivot gmath.Vec2
	MinAngle float32
	MaxAngle float32
	StartAngle float32
	Duration float32
	Cycles int
}



func (s *Swing) Run(out chan<- Step) {
	var step Step
	defer lastStep(out)

	seq := gween.NewSequence(
		gween.New(s.StartAngle, s.MinAngle, s.Duration/2, ease.InOutQuart),
		gween.New(s.MinAngle, s.MaxAngle, s.Duration, ease.InOutQuart),
		gween.New(s.MaxAngle, s.MinAngle, s.Duration, ease.InOutQuart),
		gween.New(s.MinAngle, s.StartAngle, s.Duration/2, ease.InOutQuart),
	)
	
	s.runSeq(seq, &step, out)
}

func (s *Swing) runSeq(seq *gween.Sequence, step *Step, out chan<- Step) {
	for {
		angle, _, done := seq.Update(1)
		step.Transform.Reset()
		if done {
			break
		}
		step.Transform.Translate(-s.Pivot.X, -s.Pivot.Y)
		step.Transform.Rotate(float64(angle))
		step.Transform.Translate(s.Pivot.Unwrap())
		out <- *step
	}
}
