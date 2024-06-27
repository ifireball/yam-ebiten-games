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
	dragging 	bool
	draggedLocation int
	draggedOffset ebiten.GeoM
}

func (f *Fruit) Update(screen *ebiten.Image) error {
	if !f.initialized {
		err := f.images.Load()
		if err != nil {
			return err
		}
		f.initialized = true
	}

	if f.dragging {
		x, y := ebiten.CursorPosition()
		ox, oy := f.draggedOffset.Apply(float64(x), float64(y))
		f.setLocation(f.draggedLocation, int(ox), int(oy))
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if loc, ofs, found := f.findLocationAt(ebiten.CursorPosition()); found {
			f.dragging = true
			f.draggedLocation = loc
			f.draggedOffset = ofs
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if f.dragging {
			f.dragging = false
			f.draggedLocation = -1
		} else {
			f.addLocation(ebiten.CursorPosition())
		}
	}

	return nil
}

func (f *Fruit) addLocation(x, y int) {
	var l Location
	l.position.Translate(float64(x), float64(y))
	l.position.Translate(-float64(fruit.Width)/2, -float64(fruit.Height)/2)
	f.locations = append(f.locations, l)
}

func (f *Fruit) findLocationAt(x, y int) (int, ebiten.GeoM, bool) {
	for i, loc := range f.locations {
		ax, ay := loc.position.Apply(-float64(x), -float64(y))
		if -fruit.Width <= ax && ax <= 0 && -fruit.Height <= ay && ay <= 0 {
			ofs := loc.position
			ofs.Translate(-float64(x), -float64(y))
			return i, ofs, true
		}
	}
	return -1, ebiten.GeoM{}, false
}

func (f *Fruit) setLocation(locIdx, x, y int) {
	var geom ebiten.GeoM
	geom.Translate(float64(x), float64(y))
	f.locations[locIdx].position = geom
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	dio := ebiten.DrawImageOptions{}
	for i := 0; i < len(f.locations); i++ {
		dio.GeoM = f.locations[i].position
		screen.DrawImage(f.images[f.locations[i].kind], &dio)
	}
}
