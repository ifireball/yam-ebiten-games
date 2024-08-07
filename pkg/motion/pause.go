package motion

import "github.com/hajimehoshi/ebiten/v2"

type Pause int

func (p Pause) Run() StepFunc {
	i := -1
	return func(step *ebiten.GeoM) bool { 
		step.Reset()
		i++
		return i < int(p)
	}
}
