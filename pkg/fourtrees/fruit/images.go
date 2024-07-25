package fruit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	Kinds = 4

	Width  = 25 * 3 * 3 / 4
	Height = 25 * 3 * 3 / 4
)

var KindNames = [Kinds]string{"orange", "lemon", "apple", "pear"}

var (
	Center = gmath.Vec2{X: Width / 2, Y: Height / 2}
)

type Images [Kinds]*ebiten.Image

func (fi *Images) Load() (err error) {
	for i := 0; i < Kinds; i++ {
		fi[i], err = resources.EbitenImageFromSVG(KindNames[i], Width, Height)
		if err != nil {
			return err
		}
	}
	return nil
}
