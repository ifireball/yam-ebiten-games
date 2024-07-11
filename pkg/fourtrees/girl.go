package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	girlWidth  = 100 * 2 * 3 / 4
	girlHeight = 150 * 2 * 3 / 4

	girlSpeed = 10

	minPosition = girlWidth/2
	maxPosition = gdata.ScreenWidth - girlWidth/2
)

type Girl struct {
	position int
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

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyF):
		g.position = g.position - girlSpeed
	case ebiten.IsKeyPressed(ebiten.KeyJ):
		g.position = g.position + girlSpeed
	}
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
	bottom := gdata.ScreenHeight * 17 / 20
	top := bottom - girlHeight
	left := middle - girlWidth/2
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(float64(left), float64(top))
	screen.DrawImage(g.sprite, &dio)
}
