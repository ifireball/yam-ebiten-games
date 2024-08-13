package fourtrees

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/resources"
)

// Shared state between all fourtrees scenes
type GState struct {
	Background *ebiten.Image
	Girl       Girl
	Fruit      Fruit
}

func (g *GState) Load() {
	var err error

	if err = resources.DecodeData("fourtrees", g); err != nil {
		panic(err)
	}
	if g.Background, err = resources.EbitenImageFromSVG("four_trees", gdata.ScreenWidth, gdata.ScreenHeight); err != nil {
		panic(err)
	}
	g.Fruit.Load()
	g.Girl.Load()
}

func (g *GState) Update() error {
	if err := g.Girl.Update(); err != nil {
		return err
	}
	if err := g.Fruit.Update(); err != nil {
		return err
	}
	g.detectFruitWin()

	return nil
}

func (g *GState) detectFruitWin() {
	var basketRect, fruitRect gmath.Rect
	g.Girl.GetBasketRect(&basketRect)
	g.Fruit.GetActiveRect(&fruitRect)
	if fruitRect.Overlap(&basketRect) {
		fmt.Printf("Caught Fruit!\n")
		g.Fruit.SetActiveWin()
	}
}

func (g *GState) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
	g.Girl.Draw(screen)
	g.Fruit.Draw(screen)
}
