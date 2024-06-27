package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/pkg/author"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	screenWidth  = 1920 * 3 / 4
	screenHeight = 1080 * 3 / 4
)

type Game struct {
	Background *ebiten.Image
	Fruit author.Fruit
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update(screen *ebiten.Image) error {
	var err error
	if g.Background == nil {
		g.Background, err = resources.EbitenImageFromSVG("four_trees", screenWidth, screenHeight)
		if err != nil {
			return err
		}
	}
	return g.Fruit.Update(screen)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
	g.Fruit.Draw(screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game authoring")

	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
