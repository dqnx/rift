package main

import "gitlab.com/rauko1753/gorl"

type TitleScene struct {
}

func (s TitleScene) HandleInput(k gorl.Key) (Transition, Scene) {
	switch k {
	case gorl.KeyEsc:
		return Next, nil
	}
	return Stay, nil
}

func (s TitleScene) Render() {
	TermDrawText(3, 1, "Rift")
}
