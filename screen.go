package main

import (
	"errors"

	dim "github.com/sean-brock/dimension"
)

type TileList map[dim.Vec]Tile

type Screen struct {
	size  dim.Vec
	pos   dim.Vec
	tiles []Tile
}


func (s *Screen) Size() dim.Vec {
	return s.size
}
func (s *Screen) Pos() dim.Vec {
	return s.pos
}

func (s *Screen) Tiles(o dim.Vec) TileList {
	var tiles TileList
	for i := range s.tiles {
		// Ignore blank tiles
		if s.tiles[i] != Blank {
			var pos dim.Vec
			pos.X = i%s.size.X + o.X + s.pos.X
			pos.Y = i/s.size.X + o.Y + s.pos.Y
			tiles[pos] = s.tiles[i]
		}
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

// MakeScreen returns a screen of the specified pos and size, with blank elements.
func MakeScreen(s, p dim.Vec) *Screen {
	t := make([]Tile, s.Dot())
	for i := range t {
		t[i] = Blank
	}
	return &Screen{size: s, pos: p, tiles: t}
}

type Outline int
const (
	Single Outline = iota
	Double 
)
// MakeScreenOutline returns a screen with lined outer edges.
func MakeScreenOutline(s dim.Vec, p dim.Vec, o Outline) *Screen {
	t := make([]Tile, s.Dot())
	for y := range s.Y {
		col := y*s.X
		for x := range s.X {
			i := x + col
			if x == 0 && y == 0 {
				if o == Double {
					t[i] = W2up2right
				} else {
					t[i] = W1up1right
				}
			} else if x == s.X && y == 0 {
				if o == Double {
					t[i] = W2up2left
				} else {
					t[i] = W1up1left
				}
			} else if x == 0 && y == s.Y {
				if o == Double {
					t[i] = W2down2right
				} else {
					t[i] = W1down1right
				}
			} else if x == s.X && y == s.Y {
				if o == Double {
					t[i] = W2down2left
				} else {
					t[i] = W1down1left
				}
			} else if x == 0 || x == s.X {
				if o == Double {
					t[i] = W2v
				} else {
					t[i] = W1v
				}
			} else if y == 0 || y == s.Y {
				if o == Double {
					t[i] = W2h
				} else {
					t[i] = W1h
				}
			} else {
				t[i] = Blank
			}
		}
	}
	return &Screen{size: s, pos: p, tiles: t}	
}