package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/ifireball/yam-ebiten-games/pkg/scenes"
	"github.com/joelschutz/stagehand"
)

// Minigame is both a scene and a scene director for the various fourtrees
// scenes and a scene manage by the program's top-leve scene director
type Minigame struct {
	scenes.Boiler
	scnDir *stagehand.SceneDirector[*GState]
}

func (mg *Minigame) Load(gState *gstate.GState, gScnDir stagehand.SceneController[*gstate.GState]) {
	mg.Boiler.Load(gState, gScnDir)

	state := &GState{}
	game := &Game{}
	demo := &Demo{}
	exit := &scenes.MiniExit[GState]{GScnDir: mg.ScnDir}

	var rules = map[stagehand.Scene[*GState]][]stagehand.Directive[*GState]{
		demo: {
			{Trigger: scenes.Exit, Dest: exit},
			{Trigger: scenes.Enter, Dest: game},
		},
		game: {
			{Trigger: scenes.Exit, Dest: demo},
		},
	}

	mg.scnDir = stagehand.NewSceneDirector[*GState](demo, state, rules)
}

func (mg *Minigame) Update() error {
	if err := mg.scnDir.Update(); err != nil {
		return err
	}
	return nil
}

func (mg *Minigame) Draw(screem *ebiten.Image) {
	mg.scnDir.Draw(screem)
}
