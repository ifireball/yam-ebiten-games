package fruit

import (
	"errors"

	"github.com/ifireball/yam-ebiten-games/pkg/sound"
)

type Sounds struct {
	Swing, Drop, Oy, Rot, Grow sound.Note
}

func (s *Sounds) Load() error {
	return errors.Join(
		loadInto(&s.Swing, "swing"),
		loadInto(&s.Drop, "drop"),
		loadInto(&s.Oy, "oy"),
		loadInto(&s.Rot, "rot"),
		loadInto(&s.Grow, "grow"),
	)
}

func loadInto(note *sound.Note, resFile string) (err error) {
	*note, err = sound.NoteFromRes(resFile)
	return
}
