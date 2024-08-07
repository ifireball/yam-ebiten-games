package scenes

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gerrors"
)

// Scene for exiting the whole game
type ExitScene struct {
	Boiler
}

func (s *ExitScene) Update() error {
	fmt.Printf("Exit scene invoked!")
	return gerrors.ErrExitGame
}

func (s *ExitScene) Draw(*ebiten.Image) {}
