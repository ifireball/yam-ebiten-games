package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/joelschutz/stagehand"
)

// stagehand Boilerplate common to all Scenes
type Boiler struct {
	ScnDir *stagehand.SceneDirector[*gstate.GState]
	GState *gstate.GState
}

func (sb *Boiler) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gdata.ScreenWidth, gdata.ScreenHeight
}

func (sb *Boiler) Load(state *gstate.GState, scnDir stagehand.SceneController[*gstate.GState]) {
	sb.ScnDir = scnDir.(*stagehand.SceneDirector[*gstate.GState])
	sb.GState = state
}

func (sb *Boiler) Unload() *gstate.GState {
	return sb.GState
}

func (sb *Boiler) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyEscape)  || keyboard.WobeeBlueGreen.IsJustReleased() {
		sb.ScnDir.ProcessTrigger(Exit)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	return nil
}
