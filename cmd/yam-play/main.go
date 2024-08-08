package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gerrors"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/ifireball/yam-ebiten-games/pkg/scenes"
	"github.com/joelschutz/stagehand"
)

func main() {
	ebiten.SetWindowSize(gdata.ScreenWidth, gdata.ScreenHeight)
	ebiten.SetWindowTitle("Yam Play")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	state := &gstate.GState{}
	fourtreesGame := &fourtrees.Game{}
	fourtreesDemo := &fourtrees.Demo{}
	exit := &scenes.ExitScene{}

	var rules = map[stagehand.Scene[*gstate.GState]][]stagehand.Directive[*gstate.GState]{
		fourtreesDemo: {
			{Trigger: scenes.Exit, Dest: exit},
			{Trigger: scenes.Enter, Dest: fourtreesGame},
		},
		fourtreesGame: {
			{Trigger: scenes.Exit, Dest: fourtreesDemo},
		},
	}

	scnMgr := stagehand.NewSceneDirector[*gstate.GState](fourtreesDemo, state, rules)

	if err := ebiten.RunGame(scnMgr); err != nil && !errors.Is(err, gerrors.ErrExitGame) {
		log.Fatal(err)
	}
}
