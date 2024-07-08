package fourtrees

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/motion"
	"github.com/tanema/gween/ease"
)

type Fruit struct {
	gdata.Fruit
	passive         []gdata.Location
	active          *gdata.Location
	activeMotion    motion.StepFunc
	activeTransform ebiten.GeoM
	images          fruit.Images
	initialized     bool
}

var (
	swing = motion.Swing{
		Pivot:      gmath.Vec2{X: fruit.Width / 2, Y: fruit.Height / 5},
		MinAngle:   -30 * math.Pi / 180,
		MaxAngle:   30 * math.Pi / 180,
		StartAngle: 0,
		Duration:   10,
		Cycles:     1,
	}
)

func (f *Fruit) Update() error {
	if !f.initialized {
		err := f.images.Load()
		if err != nil {
			return err
		}
		f.passive = f.Locations[:len(f.Locations)-1]
		f.active = &f.Locations[len(f.Locations)-1]
		f.activeMotion = fruitFall(f.active.Position).Run()
		f.initialized = true
	}
	if f.active != nil {
		if !f.activeMotion(&f.activeTransform) {
			f.activeTransform.Reset()
			f.makeActive(rand.Intn(len(f.passive)))
			//f.active = nil
			//f.passive = f.Locations
			f.activeMotion = fruitFall(f.active.Position).Run()
		}
	}
	return nil
}

func (f *Fruit) Draw(screen *ebiten.Image) {
	dio := ebiten.DrawImageOptions{}
	for i := range f.passive {
		dio.GeoM.Reset()
		dio.GeoM.Translate(f.Locations[i].Position.Unwrap())
		screen.DrawImage(f.images[f.Locations[i].Kind], &dio)
	}
	if f.active != nil {
		dio.GeoM.Reset()
		dio.GeoM.Concat(f.activeTransform)
		dio.GeoM.Translate(f.active.Position.Unwrap())
		screen.DrawImage(f.images[f.active.Kind], &dio)
	}
}

func fruitFall(from gmath.Vec2) motion.Motion {
	drop := motion.Trnaslate{
		Duration: 180,
		To: gmath.Vec2{Y: Ground - from.Y - fruit.Height / 2}, 
		Easing: ease.OutBounce,
	}
	grow := motion.Scale{
		Duration: 60,
		To: gmath.Vec2{X: 1, Y: 1},
		Pivot: gmath.Vec2{X: fruit.Width / 2, Y: fruit.Height / 2},
		Easing: ease.OutCubic,
	}
	return motion.Chain(&swing, &drop, &grow)
}

func (f *Fruit) makeActive(idx int) {
	// Failsafe, do noting if index does not point to a passive location
	if idx >= len(f.passive) {
		return
	}
	tmp := f.passive[idx]
	f.passive[idx] = *f.active
	f.active = &tmp
}
