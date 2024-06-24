package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	screenWidth  = 1920 * 3 / 4
	screenHeight = 1080 * 3 / 4
)

type Game struct {
	Background *ebiten.Image
	Girl fourtrees.Girl
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.Background == nil {
		image, err := resources.ImageFromSVG("four_trees", screenWidth, screenHeight)
		if err != nil {
			return err
		}
		g.Background, _ = ebiten.NewImageFromImage(image, ebiten.FilterDefault)
	}
	return g.Girl.Update(screen)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
	g.Girl.Draw(screen)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Yam Play")

	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
