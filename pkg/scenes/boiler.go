package scenes

import (
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/joelschutz/stagehand"
)

// stagehand Boilerplate common to all Scenes
type Boiler struct {
	ScnDir *stagehand.SceneDirector[*gstate.GState]
	GState *gstate.GState
}

func (sb *Boiler) Load(state *gstate.GState, scnDir stagehand.SceneController[*gstate.GState]) {
	sb.ScnDir = scnDir.(*stagehand.SceneDirector[*gstate.GState])
	sb.GState = state
}

func (sb *Boiler) Unload() *gstate.GState {
	return sb.GState
}
