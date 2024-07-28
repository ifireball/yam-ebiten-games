package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
)

type Girl struct {
	Controller girl.Controller
	position float64
	sprite   *ebiten.Image
}

func (g *Girl) Update() error {
	var err error
	if g.sprite == nil {
		g.sprite, err = resources.EbitenImageFromSVG("basket_girl", girl.Width, girl.Height)
		if err != nil {
			return err
		}
	}
	if g.position == 0 {
		g.position = gdata.ScreenWidth / 2
	}

	g.Controller.Control(&g.position)

	if g.position < girl.MinPosition {
		g.position = girl.MinPosition
	}
	if g.position > girl.MaxPosition {
		g.position = girl.MaxPosition
	}
	return nil
}

func (g *Girl) Draw(screen *ebiten.Image) {
	middle := g.position
	left := middle - girl.Width/2
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(left, girl.ScreenTop)
	screen.DrawImage(g.sprite, &dio)
}

func (g *Girl) GetBasketRect(r *gmath.Rect) {
	r.TopLeft.X = g.position - girl.BasketWidth / 2
	r.BottomRight.X = g.position + girl.BasketWidth / 2
	r.TopLeft.Y = girl.ScreenBasketTop
	r.BottomRight.Y = girl.ScreenBasketBottom
}
