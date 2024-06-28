package author

import (
	"github.com/hajimehoshi/ebiten"
)

type Mouse struct {
	GeoM ebiten.GeoM
	draggedGeoM *ebiten.GeoM
	draggedOffset ebiten.GeoM
}

func (m *Mouse) UpdateGeoM() {
	x, y := ebiten.CursorPosition()
	m.GeoM.Reset()
	m.GeoM.Translate(float64(x), float64(y))
}

func (m *Mouse) StartDrag(draggedGeoM *ebiten.GeoM) {
	m.draggedGeoM = draggedGeoM
	GeomSubstract(draggedGeoM, &m.GeoM, &m.draggedOffset)
}

func (m *Mouse) Drop() {
	m.draggedGeoM = nil
}

func (m *Mouse) UpdateDrag() {
	if !m.IsDragging() {
		return
	}
	*m.draggedGeoM = m.GeoM
	m.draggedGeoM.Concat(m.draggedOffset)
}

func (m *Mouse) IsDragging() bool {
	return m.draggedGeoM != nil
}

func (m *Mouse) IsTouchingRect(r *ebiten.GeoM, rw, rh float64) bool {
	return GeomInRect(&m.GeoM, r, rw, rh)
}
