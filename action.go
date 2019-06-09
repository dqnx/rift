package main

import dim "github.com/sean-brock/dimension"

// Action is a command pattern and returns a callback for something to do with an Actor.
type Action struct {
	actor  *Actor
	action func(*Game, *Actor) bool
}

// Perform exercises the Action function on the Actor.
func (a *Action) Perform(g *Game) bool {
	return a.action(g, a.actor)
}

// MakeWalkAction advances an actor one unit of movement.
func MakeWalkAction(a *Actor, dv dim.Vec) *Action {
	action := new(Action)
	action.actor = a
	action.action = func(g *Game, actor *Actor) bool {
		newPos := a.Position.Add(dv)
		if newPos.Inside(g.Size().Add(dim.V(-1, -1))) {
			actor.Move(dv)
			return true
		}
		return false
	}
	return action
}
