package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/joelschutz/stagehand"
)

// stagehand Boilerplate common to all Scenes
type ProtoBoiler[T any] struct {
	ScnDir *stagehand.SceneDirector[*T]
	GState *T
}

func (sb *ProtoBoiler[T]) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gdata.ScreenWidth, gdata.ScreenHeight
}

func (sb *ProtoBoiler[T]) Load(state *T, scnDir stagehand.SceneController[*T]) {
	sb.ScnDir = scnDir.(*stagehand.SceneDirector[*T])
	sb.GState = state
}

func (sb *ProtoBoiler[T]) Unload() *T {
	return sb.GState
}

func (sb *ProtoBoiler[T]) Update() error {
	keyboard.WobeeBlueGreen.Update()

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) || keyboard.WobeeBlueGreen.IsJustReleased() {
		sb.ScnDir.ProcessTrigger(Exit)
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	return nil
}
