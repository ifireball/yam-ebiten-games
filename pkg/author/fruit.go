package author

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
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

	f.mouse.UpdateLocation()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if loc := f.findLocationAtMouse(); loc != nil {
			f.mouse.StartDrag(&loc.position)
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if f.mouse.IsDragging() {
			f.mouse.Drop()
		} else {
			f.addLocationAtMouse()
		}
	}

	return nil
}

func (f *Fruit) addLocationAtMouse() {
	var l Location
	l.position = f.mouse.Location
	l.position.Sub(&gmath.Vec2{X: float64(fruit.Width)/2, Y: float64(fruit.Height)/2})
	f.locations = append(f.locations, l)
}

func (f *Fruit) findLocationAtMouse() *Location {
	for i := range f.locations {
		loc := &f.locations[i]
		if f.mouse.IsTouchingRect(&loc.position, fruit.Width, fruit.Height) {
			return loc
		}
	}
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	f.mouse.UpdateLocation()
	f.mouse.UpdateDrag()
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < len(f.locations); i++ {
		dio.GeoM.Reset()
		dio.GeoM.Translate(f.locations[i].position.Unwrap())
		screen.DrawImage(f.images[f.locations[i].kind], &dio)
	}
}
