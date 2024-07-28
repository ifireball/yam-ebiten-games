package girl

import "github.com/ifireball/yam-ebiten-games/pkg/gdata"

const (
	Width  = 100 * 2 * 3 / 4
	Height = 150 * 2 * 3 / 4

	MinPosition = Width/2
	MaxPosition = gdata.ScreenWidth - Width/2

	Speed = 10
	BasketWidth = 88 * 2 * 3 / 4

	BasketTop = 2 * 2 * 3 / 4
	BasketBottom = 20 * 2 * 3 / 4

	ScreenBottom = gdata.ScreenHeight * 17 / 20
	ScreenTop = ScreenBottom - Height

	ScreenBasketTop = ScreenTop + BasketTop
	ScreenBasketBottom = ScreenTop + BasketBottom
)
