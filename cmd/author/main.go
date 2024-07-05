package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/author"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	screenWidth  = 1920 * 3 / 4
	screenHeight = 1080 * 3 / 4
)

type Game struct {
	background *ebiten.Image
	Fruit      author.Fruit
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	var err error
	if g.background == nil {
		g.background, err = resources.EbitenImageFromSVG("four_trees", screenWidth, screenHeight)
		if err != nil {
			return err
		}
	}
	return g.Fruit.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, &ebiten.DrawImageOptions{})
	g.Fruit.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func DataFilePath() string {
	_, goFile, _, ok := runtime.Caller(0)
	if !ok {
		panic(errors.New("can't locate the data file path"))
	}
	return path.Join(path.Dir(path.Dir(path.Dir(goFile))), "resources", "data", "fourtrees.json")
}

func Load() *Game {
	var game Game
	f, err := os.Open(DataFilePath())
	if errors.Is(err, os.ErrNotExist) {
		return &game
	} else if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	if err := dec.Decode(&game); err != nil {
		panic(err)
	}
	return &game
}

func Save(game *Game) {
	if err := os.MkdirAll(path.Dir(DataFilePath()), 0755); err != nil {
		panic(err)
	}
	f, err := os.Create(DataFilePath())
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	if err := enc.Encode(game); err != nil {
		panic(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game authoring")

	game := Load()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	Save(game)
}
