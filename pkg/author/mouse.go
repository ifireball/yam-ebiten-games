package author

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type Mouse struct {
	pressLocation gmath.Vec2
	dragStarted   bool
	draggedLoc    *gmath.Vec2
	draggedOffset gmath.Vec2
	dropFunc      func()
}

type MouseUpdateHandlers struct {
	OnClick     func(mouseLoc *gmath.Vec2)
	OnDragStart func(
		mouseLoc *gmath.Vec2,
		setDraggedLoc func(draggedLoc *gmath.Vec2, onDrop func()),
	)
}

func (m *Mouse) Update(handlers MouseUpdateHandlers) {
	var location gmath.Vec2
	location.SetInts(ebiten.CursorPosition())

	switch {
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft):
		m.pressLocation = location
	case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
		if !m.dragStarted && m.pressDistance(location) > 1 {
			m.dragStarted = true
			if handlers.OnDragStart != nil {
				handlers.OnDragStart(&m.pressLocation, func(draggedLoc *gmath.Vec2, onDrop func()) {
					m.draggedLoc = draggedLoc
					m.draggedOffset = *draggedLoc
					m.draggedOffset.Sub(&location)
					m.dropFunc = onDrop
				})
			}
		}
	case inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft):
		if m.dragStarted {
			m.dragStarted = false
			m.draggedLoc = nil
			if m.dropFunc != nil {
				m.dropFunc()
			}
		} else if handlers.OnClick != nil {
			handlers.OnClick(&location)
		}
	}
}

func (m *Mouse) pressDistance(location gmath.Vec2) float64 {
	return max(
		math.Abs(location.X-m.pressLocation.X),
		math.Abs(location.Y-m.pressLocation.Y),
	)
}

func (m *Mouse) Draw() {
	if m.draggedLoc != nil {
		m.draggedLoc.SetInts(ebiten.CursorPosition())
		m.draggedLoc.Add(&m.draggedOffset)
	}
}
