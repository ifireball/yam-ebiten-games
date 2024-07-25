package fruit

import (
	"errors"

	"github.com/ifireball/yam-ebiten-games/pkg/sound"
)

type Sounds struct {
	Swing, Drop, Oy, Rot, Grow sound.Note
	KindWin [Kinds]sound.Note
}

func (s *Sounds) Load() error {
	return errors.Join(
		loadInto(&s.Swing, "swing"),
		loadInto(&s.Drop, "drop"),
		loadInto(&s.Oy, "oy"),
		loadInto(&s.Rot, "rot"),
		loadInto(&s.Grow, "grow"),
		s.loadWin(),
	)
}

func loadInto(note *sound.Note, resFile string) (err error) {
	*note, err = sound.NoteFromRes(resFile)
	return
}

func (s *Sounds) loadWin() (err error) {
	for i, n := range KindNames {
		if err = loadInto(&s.KindWin[i], n); err != nil {
			return
		}
	}
	return
}
