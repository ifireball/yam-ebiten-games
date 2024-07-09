package motion

import "github.com/hajimehoshi/ebiten/v2"

type combine []Motion

func (c combine) Run() StepFunc {
	stepFuncs := make([]StepFunc, len(c))
	for i, m := range(c) {
		stepFuncs[i] = m.Run()
	}
	return func(step *ebiten.GeoM) bool {
		var subStep ebiten.GeoM
		step.Reset()
		for _, sf := range(stepFuncs) {
			if !sf(&subStep) {
				return false
			}
			step.Concat(subStep)
		}
		return true
	}
}

func Combine(m... Motion) combine {
	return combine(m)
}
