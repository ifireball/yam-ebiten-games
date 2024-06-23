package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/resources"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
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
	
		imageSvg, err := oksvg.ReadIconStream(file)
		if err != nil {
			return err
		}
		
		imageBitmap := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
		scannerGV := rasterx.NewScannerGV(screenWidth, screenHeight, imageBitmap, imageBitmap.Bounds())
		dasher := rasterx.NewDasher(screenWidth, screenHeight, scannerGV)
		//imageSvg.Transform = imageSvg.Transform.Scale(0.5, 0.5)
		imageSvg.Draw(dasher, 1.0)
		
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