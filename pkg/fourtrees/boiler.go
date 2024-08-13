package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/scenes"
)

type Boiler struct {
	scenes.ProtoBoiler[GState]
}

func (b *Boiler) Update() error {
	if err := b.ProtoBoiler.Update(); err != nil {
		return err
	}
	if err := b.GState.Update(); err != nil {
		return err
	}
	return nil
}

func (b *Boiler) Draw(screen *ebiten.Image) {
	b.GState.Draw(screen)	
}
