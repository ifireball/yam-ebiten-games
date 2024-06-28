package author

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
)

type Location struct {
	position ebiten.GeoM
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

	f.mouse.UpdateGeoM()
	f.mouse.UpdateDrag()
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
	l.position = f.mouse.GeoM
	l.position.Translate(-float64(fruit.Width)/2, -float64(fruit.Height)/2)
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
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < len(f.locations); i++ {
		dio.GeoM = f.locations[i].position
		screen.DrawImage(f.images[f.locations[i].kind], &dio)
	}
}
