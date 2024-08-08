package fourtrees

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ifireball/yam-ebiten-games/pkg/entities/buttons"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/ifireball/yam-ebiten-games/pkg/scenes"
	"github.com/ifireball/yam-ebiten-games/resources"
	"github.com/joelschutz/stagehand"
)

type Game struct {
	scenes.Boiler
	Background *ebiten.Image
	Girl       Girl
	Fruit      Fruit
	Buttons    buttons.Colored
}

func (g *Game) Load(state *gstate.GState, scnDir stagehand.SceneController[*gstate.GState]) {
	var err error
	g.Boiler.Load(state, scnDir)
	if err = resources.DecodeData("fourtrees", g); err != nil {
		panic(err)
	}
	g.Girl.Controller = &girl.WobeeCtrl{Fruit: &g.Fruit}
	if g.Background, err = resources.EbitenImageFromSVG("four_trees", gdata.ScreenWidth, gdata.ScreenHeight); err != nil {
		panic(err)
	}
	g.Fruit.Load()
	g.Girl.Load()
	g.Buttons.Load()
}

func (g *Game) Update() error {
	//var err error
	g.Boiler.Update()
	// if g.Background == nil {
	// 	g.Background, err = resources.EbitenImageFromSVG("four_trees", gdata.ScreenWidth, gdata.ScreenHeight)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	if err := g.Girl.Update(); err != nil {
		return err
	}
	if err := g.Fruit.Update(); err != nil {
		return err
	}
	g.detectFruitWin()

	if err := g.Buttons.Update(); err != nil {
		return err
	}

	return nil
}

func (g *Game) detectFruitWin() {
	var basketRect, fruitRect gmath.Rect
	g.Girl.GetBasketRect(&basketRect)
	g.Fruit.GetActiveRect(&fruitRect)
	if fruitRect.Overlap(&basketRect) {
		fmt.Printf("Caught Fruit!\n")
		g.Fruit.SetActiveWin()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
	g.Girl.Draw(screen)
	g.Fruit.Draw(screen)
	g.Buttons.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
