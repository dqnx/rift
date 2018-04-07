package main

import "gitlab.com/rauko1753/gorl"

type CharacterScene struct {
	selected int
}

func (s *CharacterScene) HandleInput(k gorl.Key) (Transition, Scene) {
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
		case 0: // create
			return Next, nil
		case 1: // back
			return Append, &TitleScene{}
		}
	}
	return Stay, nil
}

func (s *CharacterScene) Render() {
	TermDrawText(5, 2, "Character Creation")
	TermDrawText(5, 10, "create")
	TermDrawText(5, 12, "back")
	TermDrawText(4, 10+s.selected*2, "-")
}
