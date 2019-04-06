package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	dim "github.com/sean-brock/dimension"
)

// TitleScene shows the title message and start options.
type TitleScene struct {
	selected int
	screen   Screen
	size     dim.Vec
}

func (s *TitleScene) Init(size dim.Vec) {
	s.size = size
	//s.screen = MakeScreenOutline(size, dim.V(1,1), Single);
	s.screen = *MakeScreen(size, dim.V(0, 0))
	s.screen.Set(dim.V(1, 1), '@')
	fmt.Println(s.screen)
}

// HandleInput selects the different menu options.
func (s *TitleScene) HandleInput(w *pixelgl.Window) (Transition, Scene) {
	if w.JustPressed(pixelgl.KeyEscape) {
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
		//case 0: // start
		//	return Append, &CharacterScene{}
		//case 1: // exit
		//	return Next, nil
		default:
		}
	}
	return Stay, nil
}

// ExportTiles exports the scene's graphics state to a list of Tiles.
func (s *TitleScene) ExportTiles() TileList {
	return s.screen.Tiles(dim.V(0, 0))
}
