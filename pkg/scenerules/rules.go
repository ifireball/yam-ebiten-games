package scenerules

import (
	"github.com/ifireball/yam-ebiten-games/pkg/gstate"
	"github.com/joelschutz/stagehand"
)

// Rules governing movement between game scenes
var Rules = map[stagehand.Scene[*gstate.GState]][]stagehand.Directive[*gstate.GState]{}
