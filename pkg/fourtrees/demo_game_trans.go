package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/entities/buttons"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/motion"
	"github.com/joelschutz/stagehand"
	"github.com/tanema/gween/ease"
)

type DemoGameTrans struct {
	Boiler
	stagehand.BaseTransition[*GState]
	buttons    buttons.Colored
	motionStep motion.StepFunc
	Motion     motion.Motion
}

var (
	BottomToPosition = &motion.Translate{
		From:     gmath.Vec2{Y: gdata.ScreenHeight - buttons.ButtonScreenTop},
		Duration: 60 * 1,
		Easing:   ease.OutCubic,
	}
	PositionToBottom = &motion.Translate{
		To:       gmath.Vec2{Y: gdata.ScreenHeight - buttons.ButtonScreenTop},
		Duration: 60 * 1,
		Easing:   ease.InCubic,
	}
)

func (d2g *DemoGameTrans) Start(from, to stagehand.Scene[*GState], scnDir stagehand.SceneController[*GState]) {
	d2g.BaseTransition.Start(from, to, scnDir)
	var state *GState
	if fromDemo, ok := from.(*Demo); ok {
		state = fromDemo.GState
	} else {
		state = from.(*Game).GState
	}
	d2g.Boiler.Load(state, scnDir)
	d2g.buttons.Load()
	if d2g.Motion == nil {
		d2g.Motion = BottomToPosition
	}
	d2g.motionStep = d2g.Motion.Run()
	d2g.motionStep(&d2g.buttons.Transform)
}

func (d2g *DemoGameTrans) Update() error {
	// We intentionally do not call the BaseTransition Update to avoid calling
	// Update on the from and to scenes.
	if err := d2g.Boiler.Update(); err != nil {
		return err
	}
	if !d2g.motionStep(&d2g.buttons.Transform) {
		d2g.End()
	}
	return nil
}

func (d2g *DemoGameTrans) Draw(screen *ebiten.Image) {
	d2g.Boiler.Draw(screen)
	d2g.buttons.Draw(screen)
}
