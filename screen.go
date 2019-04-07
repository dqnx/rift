package main

import (
	"errors"

	dim "github.com/sean-brock/dimension"
	"golang.org/x/text/encoding/charmap"
)

type TileList map[dim.Vec]byte

type Screen struct {
	size  dim.Vec
	pos   dim.Vec
	tiles []byte
}

func (s *Screen) Size() dim.Vec {
	return s.size
}

func (s *Screen) Pos() dim.Vec {
	return s.pos
}

func (s *Screen) Tiles(o dim.Vec) TileList {
	tiles := make(TileList)
	for i := range s.tiles {
		var pos dim.Vec
		pos.X = i%s.size.X + o.X + s.pos.X
		pos.Y = i/s.size.X + o.Y + s.pos.Y
		tiles[pos] = s.tiles[i]
	}
	return tiles
}

// Set changes a tile in the Screen to the input rune
func (s *Screen) Set(p dim.Vec, r rune) error {
	i := p.X + p.Y*s.size.X
	if i < len(s.tiles) {
		if c, ok := charmap.CodePage437.EncodeRune(r); ok {
			s.tiles[i] = c
			return nil
		} else {
			return errors.New("failed to convert rune to cp437")
		}
	}
	return errors.New("index in out of range")
}

// MakeScreen returns a screen of the specified pos and size, with blank elements.
func MakeScreen(s, p dim.Vec) *Screen {
	t := make([]byte, s.Dot())
	blank, _ := charmap.CodePage437.EncodeRune(' ')
	for i := range t {
		t[i] = blank
	}
	return &Screen{size: s, pos: p, tiles: t}
}
