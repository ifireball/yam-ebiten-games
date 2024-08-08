package keyboard

import "github.com/hajimehoshi/ebiten/v2"

const maxReleaseCycles = 5

type Combo struct {
	keys          []ebiten.Key
	releaseCycles int
	justReleased  bool
}

func NewCombo(keys ...ebiten.Key) Combo {
	return Combo{keys: keys}
}

func (c *Combo) Update() {
	numPressed := 0
	for _, key := range c.keys {
		if ebiten.IsKeyPressed(key) {
			numPressed++
		}
	}

	if c.justReleased {
		// Just released state holds for only one cycle
		c.justReleased = false
	}
	switch {
	case numPressed >= len(c.keys):
		c.releaseCycles = -maxReleaseCycles
	case numPressed <= 0:
		if c.releaseCycles < 0 {
			c.justReleased = true
			c.releaseCycles = 0
		}
	default:
		if c.releaseCycles < 0 {
			c.releaseCycles++
		}
	}
}

func (c *Combo) IsJustReleased() bool {
	return c.justReleased
}
