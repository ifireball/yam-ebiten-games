package motion

import "github.com/hajimehoshi/ebiten/v2"

var Infinity = infinity{}

type infinity struct{}

func (infinity) Run() StepFunc {
	return func(step *ebiten.GeoM) bool { 
		step.Reset()
		return true 
	}
}
