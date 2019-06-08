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
	s.title = *StringToScreen("RIFT", dim.V(4, 11))
	s.options[0] = selector{*StringToScreen("start", dim.V(4, 7)), &DungeonScene{}}
	s.options[1] = selector{*StringToScreen("options", dim.V(4, 6)), &OptionsScene{}}
	s.options[2] = selector{*StringToScreen("exit", dim.V(4, 5)), nil}
}

// HandleInput selects the different menu options.
func (s *TitleScene) HandleInput(w *pixelgl.Window) (Transition, Scene) {
	if w.JustPressed(pixelgl.KeyEscape) {
		return Next, nil
	}
	if w.JustPressed(pixelgl.KeyComma) {
		if s.selected < len(s.options)-1 {
			s.selected++
		}
	}
	if w.JustPressed(pixelgl.KeyI) {
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

func (s *TitleScene) Update(g *Game) {}

// ExportTiles exports the scene's graphics state to a list of Tiles.
func (s *TitleScene) ExportTiles() TileList {
	zero := dim.V(0, 0)
	t := s.title.Tiles(zero)
	// Combine the options with the title tiles.
	for i := range s.options {
		if i == s.selected {
			pos := s.options[i].Screen.Pos().Add(dim.V(-1, 0))
			t[pos] = RuneCP437('>')
		}
		for k, v := range s.options[i].Screen.Tiles(zero) {
			t[k] = v
		}
	}
	return t
}

func (s *TitleScene) ExportText() TextList {
	return TextList{}
}
