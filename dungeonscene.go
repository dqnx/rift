package main

import "gitlab.com/rauko1753/gorl"

const (
	mapWidth  = 60
	mapHeight = 20
)

type DungeonScene struct {
	// init stores if the scene has run Init and created the dungeon.
	init     bool
	tileGrid []*gorl.Tile
}

func (s *DungeonScene) Init() {
	if s.init {
		return
	}

	s.tileGrid = GenerateDungeon(mapWidth, mapHeight)
}

func (s *DungeonScene) HandleInput(k gorl.Key) (Transition, Scene) {
	switch k {
	case gorl.KeyEsc:
		return Append, &TitleScene{}
	}
	return Stay, nil
}

func (s *DungeonScene) Render() {
	TermDrawText(5, 2, "Dungeon")
	TermDrawTileGrid(2, 3, s.tileGrid)
}
