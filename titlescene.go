package main

import (
	"github.com/faiface/pixel/pixelgl"
	dim "github.com/sean-brock/dimension"
)

type selector struct {
	Screen
	Scene
}

// TitleScene shows the title message and start options.
type TitleScene struct {
	title    Screen
	options  [3]selector
	selected int
	size     dim.Vec
}

func (s *TitleScene) Init(size dim.Vec) {
	s.size = size
	//s.screen = MakeScreenOutline(size, dim.V(1,1), Single);
	s.title = *StringToScreen("RIFT", dim.V(4, 4))
	s.options[0] = selector{*StringToScreen("start", dim.V(4, 6)), TitleScene{}}
	s.options[1] = selector{*StringToScreen("options", dim.V(4, 7)), TitleScene{}}
	s.options[2] = selector{*StringToScreen("exit", dim.V(4, 8)), nil}
}

// HandleInput selects the different menu options.
func (s *TitleScene) HandleInput(w *pixelgl.Window) (Transition, Scene) {
	if w.JustPressed(pixelgl.KeyEscape) {
		return Next, nil
	}
	if w.JustPressed(pixelgl.KeyU) {
		if s.selected < len(s.options)-1 {
			s.selected++
		}
	}
	if w.JustPressed(pixelgl.KeyJ) {
		if s.selected > 0 {
			s.selected--
		}
	}
	if w.JustPressed(pixelgl.KeyEnter) {
		nextScene := s.options[s.selected].Scene
		if nextScene == nil {
			return Next, nil // exit
		}
		return Append, nextScene
	}
	return Stay, nil
}

// ExportTiles exports the scene's graphics state to a list of Tiles.
func (s *TitleScene) ExportTiles() TileList {
	return s.screen.Tiles(dim.V(0, 0))
}
