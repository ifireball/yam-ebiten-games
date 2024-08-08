package fourtrees

import (
	"errors"
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
	sounds			fruit.Sounds
	activeWin       bool
}

var (
	normalSize = gmath.ScaleVec2(1)

	swing = motion.Swing{
		Pivot:      gmath.Vec2{X: fruit.Width / 2, Y: fruit.Height / 5},
		MinAngle:   -30 * math.Pi / 180,
		MaxAngle:   30 * math.Pi / 180,
		StartAngle: 0,
		Duration:   10,
		Cycles:     1,
	}
	grow = motion.Scale{Duration: 60, To: normalSize, Pivot: fruit.Center, Easing: ease.OutCubic}
	win = motion.Combine(
		&motion.Scale{
			Pivot:    fruit.Center,
			From:     normalSize,
			To:       gmath.ScaleVec2(2),
			Duration: 2.5 * 60,
			Easing:   ease.OutQuart,
		},
		&motion.Trnaslate{To: gmath.Vec2{Y: -fruit.Height}, Duration: 2.5*60, Easing: ease.OutQuart},
	)	
)

func (f *Fruit) Load() {
	var err error

	if err = errors.Join(f.images.Load(), f.sounds.Load()); err != nil {
		panic(err)
	}
	f.passive = f.Locations[:len(f.Locations)-1]
	f.active = &f.Locations[len(f.Locations)-1]
	f.activeMotion = f.fruitFall(f.active.Position).Run()
}

func (f *Fruit) Update() (err error) {
	if f.active != nil {
		if !f.activeMotion(&f.activeTransform) {
			f.activeWin = false
			f.activeTransform.Reset()
			f.makeActive(rand.Intn(len(f.passive)))
			//f.active = nil
			//f.passive = f.Locations
			f.activeMotion = f.fruitFall(f.active.Position).Run()
		}
	}
	return
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

func (f *Fruit) fruitFall(from gmath.Vec2) motion.Motion {
	ground := gmath.Vec2{Y: Ground - from.Y - fruit.Height/2}
	squished := gmath.Vec2{X: 0.5, Y: 0}
	fruitBottom := gmath.Vec2{X: fruit.Width / 2, Y: fruit.Height}

	drop := motion.Trnaslate{Duration: 180, To: ground, Easing: ease.OutBounce}
	rot := motion.Combine(
		&motion.Scale{Duration: 60, From: normalSize, To: squished, Pivot: fruitBottom, Easing: ease.OutCubic},
		motion.PlaceAt(ground),
	)
	return motion.Chain(
		&f.sounds.Swing, &swing, 
		&f.sounds.Drop, 
		motion.Combine(&drop, motion.Chain(motion.Pause(60), &f.sounds.Oy, motion.Infinity)),
		&f.sounds.Rot, &rot, 
		&f.sounds.Grow, &grow,
	)
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

func (f *Fruit) GetActiveRect(r *gmath.Rect) {
	r.Zero()
	if f.active == nil || f.activeWin {
		// Return zero rect during active win animation so we don't detect the
		// same collision more then once
		return
	}
	r.TopLeft.ApplyGeoM(&f.activeTransform)
	r.TopLeft.Add(&f.active.Position)
	// This is not accurate beucase it does not take scaling and rotation in the
	// activeTransform into account, but good enough for translation, and simple
	r.BottomRight = r.TopLeft
	r.BottomRight.AddPair(fruit.Width, fruit.Height)
}

func (f *Fruit) GetActiveType() int {
	if f.active == nil || f.activeWin {
		return -1
	}
	return f.active.Kind
}

func (f *Fruit) SetActiveWin() {
	if f.active == nil || f.activeWin {
		return
	}
	var stopPosition gmath.Vec2
	stopPosition.ApplyGeoM(&f.activeTransform)
	f.activeMotion = f.fruitWin(stopPosition, f.active.Kind).Run()
	f.activeWin = true
}

func (f *Fruit) fruitWin(at gmath.Vec2, kind int) motion.Motion {
	return motion.Chain(
		&f.sounds.KindWin[kind],
		motion.Combine(win, motion.PlaceAt(at)),
		&f.sounds.Grow, 
		&grow,
	)
}
