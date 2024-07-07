package fourtrees

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/fruit"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/motion"
)

type Fruit struct {
	gdata.Fruit
	passive         []gdata.Location
	active          *gdata.Location
	activeMotion    chan motion.Step
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
		f.activeMotion = make(chan motion.Step)
		go fruitFall().Run(f.activeMotion)
		f.initialized = true
	}
	if f.active != nil {
		step := <-f.activeMotion
		if step.Stop {
			f.activeTransform.Reset()
			//f.active = nil
			//f.passive = f.Locations
			go fruitFall().Run(f.activeMotion)
		} else {
			f.activeTransform = step.Transform
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

func fruitFall() motion.Motion {
	cycle := motion.Chain(&swing, motion.Pause(60*3))
	return motion.Chain(cycle, cycle, cycle, motion.Pause(60))
}
