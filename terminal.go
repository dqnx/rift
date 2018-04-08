package main

import "gitlab.com/rauko1753/gorl"

// TermDrawText calls TermDraw multiple times to draw multiple characters.
func TermDrawText(x int, y int, s string) {
	for i, r := range s {
		gorl.TermDraw(x+i, y, gorl.Char(rune(r)))
	}
}

func TermDrawTileGrid(x int, y int, tileGrid []*gorl.Tile) {
	for _, t := range tileGrid {
		gorl.TermDraw(x+t.Offset.X, y+t.Offset.Y, t.Face)
	}
}
