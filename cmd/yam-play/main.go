package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/resources"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

const (
	screenWidth  = 1920 / 2
	screenHeight = 1080 / 2
)

type Game struct{
	Background *ebiten.Image
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.Background == nil {
		file, err := resources.Resources.Open("svg/four_trees.svg")
		if err != nil {
			return err
		}
		defer file.Close()

		cvs, err := canvas.ParseSVG(file)
		if err != nil {
			return err
		}
		imageBitmap := rasterizer.Draw(cvs, 1.0, canvas.DefaultColorSpace)

		g.Background, _ = ebiten.NewImageFromImage(imageBitmap, ebiten.FilterDefault)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Yam Play")

	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}