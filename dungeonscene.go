package main

import "gitlab.com/rauko1753/gorl"

//import "fmt"

const (
	mapWidth  = 60
	mapHeight = 20
)

type DungeonScene struct {
	// init stores if the scene has run Init and created the dungeon.
	init     bool
	tileGrid []*gorl.Tile
	player   *Actor
}

func (s *DungeonScene) Init() {
	if s.init {
		return
	}
	//fmt.Println("DungeonScene Init")

	s.tileGrid = GenerateDungeon(mapWidth, mapHeight)
	s.player = &Actor{Face: gorl.Char('@')}
	s.player.SetViewDistance(6)

	// Place player at first available tile for now.
	for _, t := range s.tileGrid {
		if t.Pass && t.Lite {
			t.Occupant = s.player
			s.player.Position = t
			break
		}
	}

	s.init = true
}

func (s *DungeonScene) HandleInput(k gorl.Key) (Transition, Scene) {
	switch k {
	case gorl.KeyEsc:
		return Append, &TitleScene{}
	}

	if move := s.MovePlayer(k); move != nil {
		s.player.Position.Handle(move)
		s.player.Handle(move)
	}

	return Stay, nil
}

func (s *DungeonScene) Render() {
	TermDrawText(5, 2, "Dungeon")

	view := gorl.Offset{2, 3}
	r := &gorl.RenderEvent{}
	f := &gorl.FoVEvent{}
	s.player.Camera.Process(f)
	for _, t := range f.FoV {
		t.Handle(r)
		gorl.TermDraw(t.Offset.X+view.X, t.Offset.Y+view.Y, r.Render)
	}
}

// KeyToEvent converts a key press into a gorl Event.
func (s *DungeonScene) MovePlayer(k gorl.Key) gorl.Event {
	// Check for movement
	if move, vector := KeyToVector(k); move {
		vector.Y *= -1 // Y axis is flipped
		// Get destination tile by getting adjacent tile from player.
		if dest := s.player.Position.Adjacent[vector]; dest != nil {
			v := &gorl.MoveEvent{Dest: dest}
			return v
		}
	}
	return nil
}
