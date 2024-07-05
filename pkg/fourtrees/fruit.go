package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
)

type Fruit struct {
	gdata.Fruit
	initialized bool
	images      fruit.Images
}

func (f *Fruit) Update() error {
	if !f.initialized {
		err := f.images.Load()
		if err != nil {
			return err
		}
		f.initialized = true
	}
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	dio := ebiten.DrawImageOptions{}
	for i := range f.Locations {
		dio.GeoM.Reset()
		dio.GeoM.Translate(f.Locations[i].Position.Unwrap())
		screen.DrawImage(f.images[f.Locations[i].Kind], &dio)
	}
}
