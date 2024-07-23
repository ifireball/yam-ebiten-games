package sound

import (
	"errors"
	"fmt"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/ifireball/yam-ebiten-games/resources"
)

const SampleRate = 44100

type Note struct {
	src io.Reader
	stream io.ReadSeeker
	player *audio.Player
}

func Context() (c *audio.Context) {
	if c = audio.CurrentContext(); c == nil {
		c = audio.NewContext(SampleRate)
	}
	return
}

func NoteFromRes(file string) (n Note, err error) {
	if n.src, err = resources.Resources.Open(fmt.Sprintf("ogg/%s.ogg", file)); err != nil {
		n.Close()
		return
	}
	if n.stream, err = vorbis.DecodeWithSampleRate(SampleRate, n.src); err != nil {
		n.Close()
		return
	}
	if n.player, err = Context().NewPlayer(n.stream); err != nil {
		n.Close()
	}
	return
}

func (n *Note) Close() error {
	var perr, terr, serr error
	if n.player != nil {
		perr = n.player.Close()
	}
	if n.stream != nil {
		if c, ok := n.stream.(io.Closer); ok {
			terr = c.Close()
		}
	}
	if n.src != nil {
		if c, ok := n.src.(io.Closer); ok {
			serr = c.Close()
		}
	}
	return errors.Join(perr, terr, serr)
}

func (n *Note) Play() (err error) {
	if err = n.player.Rewind(); err != nil {
		return
	}
	n.player.Play()
	return
}
