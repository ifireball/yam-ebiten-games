package protocols

import "github.com/ifireball/yam-ebiten-games/pkg/gmath"

type WithActiveRect interface {
	GetActiveRect(r *gmath.Rect)
}
