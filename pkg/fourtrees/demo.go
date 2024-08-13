package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/ifireball/yam-ebiten-games/pkg/scenes"
	"github.com/joelschutz/stagehand"
)

type Demo struct {
	Boiler
}

func (d *Demo) Load(state *GState, scnDir stagehand.SceneController[*GState]) {
	d.Boiler.Load(state, scnDir)
	d.GState.Girl.Controller = &girl.AutoCtrl{Fruit: &d.GState.Fruit}
}

func (d *Demo) Update() error {
	if inpututil.IsKeyJustReleased(keyboard.WobeeRed) {
		d.ScnDir.ProcessTrigger(scenes.Enter)
	}
	return d.Boiler.Update()
}
