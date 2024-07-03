package author

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type Location struct {
	Position gmath.Vec2
	Kind     int
}

type Fruit struct {
	initialized bool
	Locations   []Location
	images      fruit.Images
	mouse       Mouse
	addKind     int
}

func (f *Fruit) Update(screen *ebiten.Image) error {
	if !f.initialized {
		err := f.images.Load()
		if err != nil {
			return err
		}
		f.initialized = true
	}

	f.mouse.Update(MouseUpdateHandlers{
		OnDragStart: func(mouseLoc *gmath.Vec2, setDraggedLoc func(draggedLoc *gmath.Vec2, onDrop func())) {
			if loc := f.findLocationAt(mouseLoc); loc != nil {
				originalPos := loc.Position
				setDraggedLoc(&loc.Position, func() {
					// Prevent dragging fruit off the screen
					if !loc.Position.InImageRect(screen.Bounds()) {
						loc.Position = originalPos
					}
				})
			}
		},
		OnClick: func(mouseLoc *gmath.Vec2) {
			if loc := f.findLocationAt(mouseLoc); loc != nil {
				loc.Kind = (loc.Kind + 1) % fruit.Kinds
				f.addKind = loc.Kind
			} else {
				f.addLocationAt(mouseLoc)
			}
		},
	})

	return nil
}

func (f *Fruit) addLocationAt(position *gmath.Vec2) {
	var l Location
	l.Position = *position
	l.Position.Sub(&gmath.Vec2{X: float64(fruit.Width) / 2, Y: float64(fruit.Height) / 2})
	l.Kind = f.addKind
	f.Locations = append(f.Locations, l)
}

func (f *Fruit) findLocationAt(position *gmath.Vec2) *Location {
	for i := range f.Locations {
		loc := &f.Locations[i]
		if position.InRect(&loc.Position, fruit.Width, fruit.Height) {
			return loc
		}
	}
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	f.mouse.Draw()
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < len(f.Locations); i++ {
		dio.GeoM.Reset()
		dio.GeoM.Translate(f.Locations[i].Position.Unwrap())
		screen.DrawImage(f.images[f.Locations[i].Kind], &dio)
	}
}
