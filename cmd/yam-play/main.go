package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gerrors"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/ifireball/yam-ebiten-games/pkg/scenerules"
	"github.com/ifireball/yam-ebiten-games/resources"
	"github.com/joelschutz/stagehand"
)

func main() {
	ebiten.SetWindowSize(gdata.ScreenWidth, gdata.ScreenHeight)
	ebiten.SetWindowTitle("Yam Play")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	state := &gstate.GState{}
	game := &fourtrees.Game{}
	if err := resources.DecodeData("fourtrees", game); err != nil {
		log.Fatal(err)
	}
	scnMgr := stagehand.NewSceneDirector[*gstate.GState](game, state, scenerules.Rules)

	if err := ebiten.RunGame(scnMgr); err != nil && !errors.Is(err, gerrors.ErrExitGame) {
		log.Fatal(err)
	}
}
