package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gerrors"
	"github.com/ifireball/yam-ebiten-games/resources"
)

func main() {
	ebiten.SetWindowSize(gdata.ScreenWidth, gdata.ScreenHeight)
	ebiten.SetWindowTitle("Yam Play")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	var game fourtrees.Game
	if err := resources.DecodeData("fourtrees", &game); err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&game); err != nil && !errors.Is(err, gerrors.ErrExitGame) {
		log.Fatal(err)
	}
}
