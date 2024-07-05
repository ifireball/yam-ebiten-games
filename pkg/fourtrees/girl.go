package fourtrees

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	girlWidth  = 100 * 2 * 3 / 4
	girlHeight = 150 * 2 * 3 / 4
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
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyA):
		g.position = 0
	case ebiten.IsKeyPressed(ebiten.KeyF):
		g.position = 1
	case ebiten.IsKeyPressed(ebiten.KeyJ):
		g.position = 2
	case ebiten.IsKeyPressed(ebiten.KeySemicolon):
		g.position = 3
	}
	return nil
}

func (g *Girl) Draw(screen *ebiten.Image) {
	screenW, screenH := screen.Size()
	middle := screenW * (g.position + 1) / 5
	bottom := screenH * 17 / 20
	top := bottom - girlHeight
	left := middle - girlWidth/2
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(float64(left), float64(top))
	screen.DrawImage(g.sprite, &dio)
}
