package main

import "gitlab.com/rauko1753/gorl"

type Actor struct {
	Face     gorl.Glyph
	Position *gorl.Tile
}

func (a *Actor) Handle(v gorl.Event) {
	switch v := v.(type) {
	case *gorl.RenderEvent:
		v.Render = a.Face
	case *gorl.MoveEvent:
		if v.Dest.Pass {
			a.Position = v.Dest
		}
	}

}
