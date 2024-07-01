package author

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
)

type Location struct {
	position gmath.Vec2
	kind     int
}

type Fruit struct {
	initialized bool
	locations   []Location
	images      fruit.Images
	mouse 		Mouse
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
				originalPos := loc.position
				setDraggedLoc(&loc.position, func() {
					// Prevent dragging fruit off the screen
					if !loc.position.InImageRect(screen.Bounds()) {
						loc.position = originalPos
					}
				})
			}
		},
		OnClick: f.addLocationAt,
	})

	return nil
}

func (f *Fruit) addLocationAt(position *gmath.Vec2) {
	var l Location
	l.position = *position
	l.position.Sub(&gmath.Vec2{X: float64(fruit.Width)/2, Y: float64(fruit.Height)/2})
	f.locations = append(f.locations, l)
}

func (f *Fruit) findLocationAt(position *gmath.Vec2) *Location {
	for i := range f.locations {
		loc := &f.locations[i]
		if position.InRect(&loc.position, fruit.Width, fruit.Height) {
			return loc
		}
	}
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	f.mouse.Draw()
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < len(f.locations); i++ {
		dio.GeoM.Reset()
		dio.GeoM.Translate(f.locations[i].position.Unwrap())
		screen.DrawImage(f.images[f.locations[i].kind], &dio)
	}
}
