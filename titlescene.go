package main

import (
	"gitlab.com/faiface/pixel"
	"gitlab.com/faiface/pixel/pixelgl"
)

// TitleScene shows the title message and start options.
type TitleScene struct {
	selected int
	screen Screen
	size dim.Pos
}

func (s *TitleScene) Init(size dim.Vec) {
	s.size = size
	//s.screen = MakeScreenOutline(size, dim.V(1,1), Single);
	s.screen = MakeScreen(size, dim.V(0,0));
}

// HandleInput selects the different menu options.
func (s *TitleScene) HandleInput(w *pixelgl.Window) (Transition, Scene) {
	if w.JustPressed(pixelgl.KeyEsc) {
		return Next, nil
	}
	if w.JustPressed(pixelgl.KeyJ) {
		if s.selected < 1 {
			s.selected++
		}
	}
	if w.JustPressed(pixelgl.KeyK) {
		if s.selected > 0 {
			s.selected--
		}
	}
	if w.JustPressed(pixelgl.KeyEnter) {
		switch s.selected {
			case 0: // start
				return Append, &CharacterScene{}
			case 1: // exit
				return Next, nil
			}
	}
	return Stay, nil
}

// Render exports the scene's graphics state to a list of Tiles.
func (s *TitleScene) ExportTiles() TileList {
	return s.screen.Tiles(dim.V(0,0));
}