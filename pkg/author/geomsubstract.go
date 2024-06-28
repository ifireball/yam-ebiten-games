package author

import (
	"github.com/hajimehoshi/ebiten"
)

func GeomSubstract(a, b, result *ebiten.GeoM) {
	*result = *b
	result.Invert()
	result.Concat(*a)
}

func GeomInRect(g, r *ebiten.GeoM, rw, rh float64) bool {
	var diff ebiten.GeoM
	GeomSubstract(g, r, &diff)
	dx, dy := diff.Apply(0, 0)
	inRect :=  dx >= 0 && dx < rw && dy >= 0 && dy <= rh
	return inRect
}
