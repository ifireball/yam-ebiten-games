package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/fourtrees/girl"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/gmath"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	girlWidth  = 100 * 2 * 3 / 4
	girlHeight = 150 * 2 * 3 / 4

	girlBasketTop = 2 * 2 * 3 / 4
	girlBasketBottom = 20 * 2 * 3 / 4
	girlBasketWidth = 88 * 2 * 3 / 4

	minPosition = girlWidth/2
	maxPosition = gdata.ScreenWidth - girlWidth/2

	screenGirlBottom = gdata.ScreenHeight * 17 / 20
	screenGirlTop = screenGirlBottom - girlHeight

	screenBasketTop = screenGirlTop + girlBasketTop
	screenBasketBottom = screenGirlTop + girlBasketBottom
)

type Girl struct {
	Controller girl.Controller
	position float64
	sprite   *ebiten.Image
}

func (g *Girl) Update() error {
	var err error
	if g.sprite == nil {
		g.sprite, err = resources.EbitenImageFromSVG("basket_girl", girlWidth, girlHeight)
		if err != nil {
			return err
		}
	}
	if g.position == 0 {
		g.position = gdata.ScreenWidth / 2
	}

	g.Controller.Control(&g.position)

	if g.position < minPosition {
		g.position = minPosition
	}
	if g.position > maxPosition {
		g.position = maxPosition
	}
	return nil
}

func (g *Girl) Draw(screen *ebiten.Image) {
	middle := g.position
	left := middle - girlWidth/2
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(left, screenGirlTop)
	screen.DrawImage(g.sprite, &dio)
}

func (g *Girl) GetBasketRect(r *gmath.Rect) {
	r.TopLeft.X = g.position - girlBasketWidth / 2
	r.BottomRight.X = g.position + girlBasketWidth / 2
	r.TopLeft.Y = screenBasketTop
	r.BottomRight.Y = screenBasketBottom
}
