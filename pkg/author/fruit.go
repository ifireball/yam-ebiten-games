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
}

func (f *Fruit) Update(screen *ebiten.Image) error {
	if !f.initialized {
		err := f.images.Load()
		if err != nil {
			return err
		}
		f.initialized = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		var l Location
		l.position.Translate(float64(x), float64(y))
		l.position.Translate(-float64(fruit.Width)/2, -float64(fruit.Height)/2)
		f.locations = append(f.locations, l)
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
