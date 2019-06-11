package main

import (
	"github.com/faiface/pixel/pixelgl"
	dim "github.com/sean-brock/dimension"
	"golang.org/x/image/colornames"
)

// DungeonScene shows the title message and start options.
type DungeonScene struct {
	border      *Screen
	size        dim.Vec
	dungeonSize dim.Vec
	dungeon     *Screen
	advance     bool
	Game
}

// Init initlizes the dungeon: game and screens.
func (s *DungeonScene) Init(size dim.Vec) {
	// Create Game
	s.Game.background = colornames.Darkgrey
	s.Game.foreground = colornames.White
	s.Game.size = size.Add(dim.V(-2, -2))
	s.Game.Init()

	// Create border
	s.size = size.Add(dim.V(-1, -1))
	s.border = MakeScreen(size, dim.V(0, 0))
	s.border.Outline('#')

	// Create dungeon floor inside border
	s.dungeonSize = size.Add(dim.V(-2, -2))
	s.dungeon = MakeScreen(s.dungeonSize, dim.V(1, 1))
}

// HandleInput selects the different menu options.
func (s *DungeonScene) HandleInput(w *pixelgl.Window) (Transition, Scene) {
	if w.JustPressed(pixelgl.KeyEscape) {
		return Append, new(TitleScene)
	}
	if w.JustPressed(pixelgl.KeyEnter) {
		s.advance = true
	}
	return Stay, nil
}

// Update processes the next actors turn.
func (s *DungeonScene) Update() {
	if !s.advance {
		return
	}
	s.Process()
	s.dungeon.Clear()
	for _, a := range s.Actors() {
		s.dungeon.Set(a.Position, '@')
	}
	s.advance = false
}

// ExportTiles exports the scene's graphics state to a list of Tiles.
func (s *DungeonScene) ExportTiles() TileList {
	zero := dim.V(0, 0)
	t := s.border.Tiles(zero)
	dungeon := s.dungeon.Tiles(zero)
	for k, v := range dungeon {
		t[k] = v
	}
	return t
}

func (s *DungeonScene) ExportText() TextList {
	return TextList{}
}
