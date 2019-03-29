package main

import (
	"errors"

	dim "github.com/sean-brock/dimension"
)

type Screen struct {
	size  dim.Vec
	pos   dim.Vec
	tiles []Tile
}

// MakeScreen returns a screen of the specified pos and size, with blank elements.
func MakeScreen(s, p dim.Vec) *Screen {
	t := make([]Tile, s.Dot())
	for i := range t {
		t[i] = Blank
	}
	return &Screen{size: s, pos: p, tiles: t}
}
func (s *Screen) Size() dim.Vec {
	return s.size
}
func (s *Screen) Pos() dim.Vec {
	return s.pos
}

func (s *Screen) Tiles(o dim.Vec) map[dim.Vec]Tile {
	tiles := make(map[dim.Vec]Tile)
	for i := range s.tiles {
		var pos dim.Vec
		pos.X = i%s.size.X + o.X
		pos.Y = i/s.size.X + o.Y
		tiles[pos] = s.tiles[i]
	}
	return tiles
}

func (s *Screen) Set(p dim.Vec, t Tile) error {
	i := p.X + p.Y*s.size.X
	if i < len(s.tiles) {
		s.tiles[i] = t
		return nil
	}
	return errors.New("Index in out of range.")
}
