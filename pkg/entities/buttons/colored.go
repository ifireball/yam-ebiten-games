package buttons

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ifireball/yam-ebiten-games/pkg/gdata"
	"github.com/ifireball/yam-ebiten-games/pkg/keyboard"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const (
	ButtonWidth  = 50 * 2 * 3 / 4
	ButtonHeight = 50 * 2 * 3 / 4

	ButtonScreenTop = gdata.ScreenHeight * 37 / 40 - ButtonHeight / 2
)

type Colored struct {
	images      [keyboard.WobeeKeysAmount * 2]*ebiten.Image
	pressed     [keyboard.WobeeKeysAmount]int8
	Transform	ebiten.GeoM
}

func (c *Colored) Load() {
	var err error

	for i, cname := range keyboard.WobeeColorNames {
		prsName := fmt.Sprintf("%s_btn_pressed", cname)
		btnName := prsName[0:len(prsName)-8]
		c.images[i*2], err = resources.EbitenImageFromSVG(btnName, ButtonWidth, ButtonHeight)
		if err != nil {
			panic(err)
		}
		c.images[i*2+1], err = resources.EbitenImageFromSVG(prsName, ButtonWidth, ButtonHeight)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Colored) Update() (err error) {
	for i, key := range keyboard.WobeeOrder {
		if ebiten.IsKeyPressed(key) {
			c.pressed[i] = 1
		} else {
			c.pressed[i] = 0
		}
	}
	return
}

func (c *Colored) Draw(screen *ebiten.Image) {
	dio := ebiten.DrawImageOptions{}
	for i, pressed := range c.pressed {
		middle := gdata.ScreenWidth * (float64(i) + 1) / (keyboard.WobeeKeysAmount + 1)
		left := middle - ButtonWidth/2
		dio.GeoM = c.Transform
		dio.GeoM.Translate(
			left,
			ButtonScreenTop,
		)
		screen.DrawImage(c.images[i*2 + int(pressed)], &dio)
	}
}
