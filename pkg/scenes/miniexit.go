package scenes

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/joelschutz/stagehand"
)

// Mini exit is a prototype scene for exiting minigames (Scene directors who
// are scenes themselves in larger games)
type MiniExit[T any] struct {
	ProtoBoiler[T]
	// The "global" scene controller that runs the larger game
	GScnDir *stagehand.SceneDirector[*gstate.GState]
}

func (me *MiniExit[T]) Update() error {
	fmt.Println("Mini exit scene invoked!")
	me.GScnDir.ProcessTrigger(Exit)
	return nil
}

func (me *MiniExit[T]) Draw(*ebiten.Image) {}
