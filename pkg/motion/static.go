package motion

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type Static struct {
	ebiten.GeoM
} 

func (s *Static) Run() StepFunc {
	return func(step *ebiten.GeoM) bool {
		*step = s.GeoM
		return true
	}
}

func PlaceAt(loc gmath.Vec2) *Static {
	var static Static
	static.Translate(loc.Unwrap())
	return &static
}
