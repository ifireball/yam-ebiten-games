package fruit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	Kinds = 4

	Width  = 25 * 3 * 3 / 4
	Height = 25 * 3 * 3 / 4
)

var imageNames = [Kinds]string{"orange", "lemon", "apple", "pear"}

type Images [Kinds]*ebiten.Image

func (fi *Images) Load() (err error) {
	for i := 0; i < Kinds; i++ {
		fi[i], err = resources.EbitenImageFromSVG(imageNames[i], Width, Height)
		if err != nil {
			return err
		}
	}
	return nil
}
