package main

import (
	//"fmt"
	"gitlab.com/rauko1753/gorl"
)

const (
	roomMaxSize = 11
	roomMinSize = 4
	maxRooms    = 16
)

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

// Center returns the center point of a Rect
func (r Rect) Center() (x, y int) {
	x = (r.X1 + r.X2) / 2
	y = (r.Y1 + r.Y2) / 2
	return
}

// Intersect checks if 2 Rects overlap
func (r1 Rect) Intersect(r2 Rect) bool {
	return r1.X1 <= r2.X2 && r1.X2 >= r2.X1 && r1.Y1 <= r2.Y2 && r1.Y2 >= r2.Y1
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
		for j := 0; j < y; j++ {
			t[i][j].Pass = false
			t[i][j].Lite = false
		}
	}

	return MapTmpl{Tiles: t}
}

func (m *MapTmpl) createRoom(r Rect) {
	for x := r.X1 + 1; x <= r.X2; x++ {
		for y := r.Y1 + 1; y <= r.Y2; y++ {
			m.Tiles[x][y].Pass = true
			m.Tiles[x][y].Lite = true
		}
	}
}

func (m *MapTmpl) createHTunnel(x1, x2, y int) {
	if x1 > x2 {
		x2, x1 = x1, x2 // swap x1 and x2
	}

	for x := x1; x <= x2; x++ {
		m.Tiles[x][y].Pass = true
		m.Tiles[x][y].Lite = true
	}
}

func (m *MapTmpl) createVTunnel(y1, y2, x int) {
	if y1 > y2 {
		y2, y1 = y1, y2 // swap y1 and y2
	}

	for y := y1; y <= y2; y++ {
		m.Tiles[x][y].Pass = true
		m.Tiles[x][y].Lite = true
	}
}

// GenerateDungeon creates a classic dungeon layout of rooms and hallways.
func GenerateDungeon(cols, rows int) []*gorl.Tile {
	dungeon := NewMapTmpl(cols, rows)
	//dungeon.createRoom(NewRect(2, 2, 5, 5))

	var rooms []Rect

	for r := 0; r < maxRooms; r++ {

		// Random width and height
		w := gorl.RandRange(roomMinSize, roomMaxSize)
		h := gorl.RandRange(roomMinSize, roomMaxSize)
		// Random within map boundaries
		x := gorl.RandRange(0, mapWidth-w-1)
		y := gorl.RandRange(0, mapHeight-h-1)

		newRoom := NewRect(x, y, w, h)
		intersect := false // flag for if the generated room intersects with existing rooms

		for _, otherRoom := range rooms {
			if newRoom.Intersect(otherRoom) {
				intersect = true
				break
			}
		}
		if !intersect {
			dungeon.createRoom(newRoom)
			newX, newY := newRoom.Center()
			if len(rooms) > 0 {
				prevX, prevY := rooms[len(rooms)-1].Center()
				// flip a coin
				if gorl.RandBool() {
					dungeon.createHTunnel(prevX, newX, prevY)
					dungeon.createVTunnel(prevY, newY, newX)
				} else {
					dungeon.createVTunnel(prevY, newY, prevX)
					dungeon.createHTunnel(prevX, newX, newY)
				}
			}
			rooms = append(rooms, newRoom)
		}
	}

	var tileMapper = func(o gorl.Offset) *gorl.Tile {
		tile := gorl.NewTile(o)
		tile.Pass = dungeon.At(o).Pass
		tile.Lite = dungeon.At(o).Lite
		return tile
	}

	tileGrid := gorl.GenTileGrid(cols, rows, tileMapper)

	return tileGrid
}
