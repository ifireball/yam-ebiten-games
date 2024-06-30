package author

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type Mouse struct {
	Location gmath.Vec2
	draggedLoc *gmath.Vec2
	draggedOffset gmath.Vec2
}

func (m *Mouse) UpdateLocation() {
	m.Location.SetInts(ebiten.CursorPosition())
}

func (m *Mouse) StartDrag(draggedLoc *gmath.Vec2) {
	m.draggedLoc = draggedLoc
	m.draggedOffset = *draggedLoc
	m.draggedOffset.Sub(&m.Location)
}

func (m *Mouse) Drop() {
	m.draggedLoc = nil
}

func (m *Mouse) UpdateDrag() {
	if !m.IsDragging() {
		return
	}
	*m.draggedLoc = m.Location
	m.draggedLoc.Add(&m.draggedOffset)
}

func (m *Mouse) IsDragging() bool {
	return m.draggedLoc != nil
}

func (m *Mouse) IsTouchingRect(r *gmath.Vec2, rw, rh float64) bool {
	return m.Location.InRect(r, rw, rh)
}
