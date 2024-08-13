package fourtrees

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ifireball/yam-ebiten-games/pkg/entities/buttons"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/joelschutz/stagehand"
)

type Game struct {
	Boiler
	Buttons    buttons.Colored
}

func (g *Game) Load(state *GState, scnDir stagehand.SceneController[*GState]) {
	g.Boiler.Load(state, scnDir)
	g.GState.Girl.Controller = &girl.WobeeCtrl{Fruit: &g.GState.Fruit}
	g.Buttons.Load()
}

func (g *Game) Update() error {
	if err := g.Boiler.Update(); err != nil {
		return err
	}
	if err := g.Buttons.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Boiler.Draw(screen)
	g.Buttons.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
