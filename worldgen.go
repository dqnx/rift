package main

import "gitlab.com/rauko1753/gorl"

// TileTmpl is a template tile used for map generation.
type TileTmpl struct {
	Face   gorl.Glyph
	Pass   bool
	Lite   bool
	Offset gorl.Offset
}

// MapTmpl is a template 2d slice used for map generation.
type MapTmpl struct {
	Tiles [][]TileTmpl
}

type rect struct {
	X1, Y1, X2, Y2 int
}

func NewMapTmpl(x, y int) MapTmpl {
	m := make([][]TileTmpl, y)
	for i := range m {
		m[i] = make([]TileTmpl, x)
	}
	return MapTmpl{Tiles: m}
}

func (m *MapTmpl) createRoom(r rect) {

}

// GenerateDungeon creates a classic dungeon layout of rooms and hallways.
//func GenerateDungeon(cols, rows int) []*gorl.Tile {
//var tileMapper = func(o gorl.Offset) *gorl.Tile {

//}
//}
