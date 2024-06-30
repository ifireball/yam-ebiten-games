package gmath

import (
	"github.com/hajimehoshi/ebiten"
)

type Vec2 struct{
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
	v.X, v.Y = v.X + v2.X, v.Y + v2.Y
}

func (v *Vec2) Sub(v2 *Vec2) {
	v.X, v.Y = v.X - v2.X, v.Y - v2.Y
}

func (v *Vec2) Neg() {
	v.X, v.Y = -v.X, -v.Y
}

func (v *Vec2) ApplyGeoM(g *ebiten.GeoM) {
	v.X, v.Y = g.Apply(v.X, v.Y)
}

func (v *Vec2) InRect(topLeft *Vec2, w, h float64) bool {
	return v.X >= topLeft.X && v.X < topLeft.X + w &&
		v.Y >= topLeft.Y && v.Y < topLeft.Y + h
}
