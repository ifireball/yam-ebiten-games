package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/entities/buttons"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/ifireball/yam-ebiten-games/resources"
)

var (
	ErrExitGame = errors.New("exit game")
)

type Game struct {
	Background *ebiten.Image
	Girl fourtrees.Girl
	Fruit fourtrees.Fruit
	Buttons buttons.Colored
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gdata.ScreenWidth, gdata.ScreenHeight
}

func (g *Game) Update() error {
	var err error
	if g.Background == nil {
		g.Background, err = resources.EbitenImageFromSVG("four_trees", gdata.ScreenWidth, gdata.ScreenHeight)
		if err != nil {
			return err
		}
	}
	if g.Girl.Controller == nil {
		g.Girl.Controller = &girl.WobeeCtrl{Fruit: &g.Fruit}
	}

	keyboard.WobeeBlueGreen.Update()

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape)  || keyboard.WobeeBlueGreen.IsJustReleased() {
		return ErrExitGame
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

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
		fmt.Printf("Caught Fruit!")
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

func main() {
	ebiten.SetWindowSize(gdata.ScreenWidth, gdata.ScreenHeight)
	ebiten.SetWindowTitle("Yam Play")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	var game Game
	if err := resources.DecodeData("fourtrees", &game); err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&game); err != nil && !errors.Is(err, ErrExitGame) {
		log.Fatal(err)
	}
}
