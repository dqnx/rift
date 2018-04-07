package main

import "gitlab.com/rauko1753/gorl"

// TitleScene shows the title message and start options.
type TitleScene struct {
	selected int
}

// HandleInput selects the different menu options.
func (s *TitleScene) HandleInput(k gorl.Key) (Transition, Scene) {
	switch k {
	case gorl.KeyEsc:
		return Next, nil
	case 'k':
		if s.selected < 1 {
			s.selected++
		}
	case 'j':
		if s.selected > 0 {
			s.selected--
		}
	case gorl.KeyEnter:
		switch s.selected {
		case 0: // start
			return Append, &CharacterScene{}
		case 1: // exit
			return Next, nil
		}
	}
	return Stay, nil
}

// Render prints the title screen.
func (s *TitleScene) Render() {
	TermDrawText(5, 2, "RIFT")
	TermDrawText(5, 10, "start")
	TermDrawText(5, 12, "exit")
	TermDrawText(4, 10+s.selected*2, "-")
}
