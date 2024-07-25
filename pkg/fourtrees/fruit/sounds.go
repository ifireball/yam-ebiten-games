package fruit

import (
	"errors"

	"github.com/ifireball/yam-ebiten-games/pkg/sound"
)

type Sounds struct {
	Drop, Grow sound.Note
}

func (s *Sounds) Load() error {
	return errors.Join(
		loadInto(&s.Drop, "drop"),
		loadInto(&s.Grow, "grow"),
	)
}

func loadInto(note *sound.Note, resFile string) (err error) {
	*note, err = sound.NoteFromRes(resFile)
	return
}
