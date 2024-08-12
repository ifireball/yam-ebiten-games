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

	exit := &scenes.MiniExit[GState]{GScnDir: mg.ScnDir}
	rules, first := mg.setupRules(exit)

	state := mg.loadState()
	mg.scnDir = stagehand.NewSceneDirector[*GState](first, state, rules)
}

func (mg *Minigame) loadState() *GState {
	state := &GState{}
	return state	
}

func (mg *Minigame) setupRules(exit stagehand.Scene[*GState]) (
	map[stagehand.Scene[*GState]][]stagehand.Directive[*GState],
	stagehand.Scene[*GState],
) {
	game := &Game{}
	demo := &Demo{}

	var rules = map[stagehand.Scene[*GState]][]stagehand.Directive[*GState]{
		demo: {
			{Trigger: scenes.Exit, Dest: exit},
			{Trigger: scenes.Enter, Dest: game},
		},
		game: {
			{Trigger: scenes.Exit, Dest: demo},
		},
	}
	return rules, demo
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
