package motion

import "github.com/hajimehoshi/ebiten/v2"

type Step struct {
	Transform ebiten.GeoM
	Stop bool
}

type Motion interface {
	Run(chan<- Step)
}
