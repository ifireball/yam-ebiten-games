package motion

import "github.com/hajimehoshi/ebiten/v2"

type StepFunc func(*ebiten.GeoM) bool

type Motion interface {
	Run() StepFunc
}
