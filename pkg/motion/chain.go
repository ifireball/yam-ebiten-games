package motion

import "github.com/hajimehoshi/ebiten/v2"

type chain []Motion

func (c chain) Run() StepFunc {
	if len(c) <= 0 {
		return func(*ebiten.GeoM) bool { return false }
	}

	motionI := 0
	motionFunc := c[motionI].Run()

	return func(step *ebiten.GeoM) bool {
		for {
			if motionFunc(step) {
				return true
			}
			if motionI++; motionI >= len(c) {
				return false
			}
			motionFunc = c[motionI].Run()
		}
	}
}

func Chain(m... Motion) chain {
	return chain(m)
}
