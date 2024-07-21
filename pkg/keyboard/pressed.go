package keyboard

import "github.com/hajimehoshi/ebiten/v2"

func IsPressedOneOf(keys... ebiten.Key) bool {
	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}
	return false
}
