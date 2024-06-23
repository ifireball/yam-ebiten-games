package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	screenWidth  = 1920 / 2
	screenHeight = 1080 / 2
)

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Yam Play")

	file, err := resources.Resources.Open("svg/four_trees.svg")
	if err != nil {
		log.Fatal(err)
	}
	buf := [1024]byte{}
	if amount, err := file.Read(buf[:]); err == nil {
		fmt.Println(string(buf[:amount]))
	}

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}