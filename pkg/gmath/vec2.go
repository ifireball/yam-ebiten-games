package gmath

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) SetInts(x, y int) {
	v.X, v.Y = float64(x), float64(y)
}

func (v *Vec2) Unwrap() (float64, float64) {
	return v.X, v.Y
}

func (v *Vec2) Reset() {
	v.X, v.Y = 0, 0
}

func (v *Vec2) Add(v2 *Vec2) {
	v.X, v.Y = v.X+v2.X, v.Y+v2.Y
}

func (v *Vec2) Sub(v2 *Vec2) {
	v.X, v.Y = v.X-v2.X, v.Y-v2.Y
}

func (v *Vec2) Neg() {
	v.X, v.Y = -v.X, -v.Y
}

func (v *Vec2) ApplyGeoM(g *ebiten.GeoM) {
	v.X, v.Y = g.Apply(v.X, v.Y)
}

func (v *Vec2) InBounds(l, t, b, r float64) bool {
	return v.X >= l && v.X < r && v.Y >= t && v.Y < b
}

func (v *Vec2) InRect(l, t, w, h float64) bool {
	return v.InBounds(l, t, l+w, t+h)
}

func (v *Vec2) InVecRect(topLeft *Vec2, w, h float64) bool {
	return v.InRect(topLeft.X, topLeft.Y, w, h)
}

func (v *Vec2) InImageRect(r image.Rectangle) bool {
	return v.InBounds(float64(r.Min.X), float64(r.Min.Y), float64(r.Max.X), float64(r.Max.Y))
}
