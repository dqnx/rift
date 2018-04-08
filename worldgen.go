package main

import "gitlab.com/rauko1753/gorl"

// TileTmpl is a template tile used for map generation.
type TileTmpl struct {
	Face gorl.Glyph
	Pass bool
	Lite bool
}

// Rect is a 2d polygon with 2 points.
type Rect struct {
	X1, Y1, X2, Y2 int
}

// NewRect creates a Rect from coordinates and size.
func NewRect(x, y, w, h int) Rect {
	return Rect{X1: x, X2: x + w, Y1: y, Y2: y + h}
}

// MapTmpl is a template 2d slice used for map generation.
type MapTmpl struct {
	Tiles [][]TileTmpl
}

// At retrieves a TileTmpl at the input offset.
func (m *MapTmpl) At(o gorl.Offset) *TileTmpl {
	return &m.Tiles[o.X][o.Y]
}

// NewMapTmpl creatsr a MapTmpl from a size.
func NewMapTmpl(x, y int) MapTmpl {
	t := make([][]TileTmpl, x)
	for i := 0; i < x; i++ {
		t[i] = make([]TileTmpl, y)
	}

	return MapTmpl{Tiles: t}
}

func (m *MapTmpl) createRoom(r Rect) {
	for x := r.X1; x <= r.X2; x++ {
		for y := r.Y1; y <= r.Y2; y++ {
			m.Tiles[x][y].Pass = true
			m.Tiles[x][y].Lite = true
		}
	}
}

// GenerateDungeon creates a classic dungeon layout of rooms and hallways.
func GenerateDungeon(cols, rows int) []*gorl.Tile {
	dungeon := NewMapTmpl(cols, rows)
	dungeon.createRoom(Rect{2, 2, 8, 5})

	var tileMapper = func(o gorl.Offset) *gorl.Tile {
		tile := gorl.NewTile(o)
		tile.Pass = dungeon.At(o).Pass
		tile.Lite = dungeon.At(o).Lite
		//tile.Face = dungeon.At(o).Face
		return tile
	}

	tileGrid := gorl.GenTileGrid(cols, rows, tileMapper)

	return tileGrid
}
