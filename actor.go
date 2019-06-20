package main

import (
	"fmt"
	"math/rand"

	dim "github.com/sean-brock/dimension"
)

// Actor is any game character or sentient creature.
type Actor struct {
	Name string
	ID
	Position dim.Vec
	// Flyweight these
	Speed int
}

// Update will return an action from an actor.
func (a *Actor) Update() {
	fmt.Println("Updated Actor:", a.ID)

}

// Move changes an Actor's position.
func (a *Actor) Move(dv dim.Vec) {
	a.Position = a.Position.Add(dv)
}

// Action returns the Actors next Action.
func (a *Actor) Action() *Action {
	r := rand.Intn(4)
	var dv dim.Vec
	switch r {
	case 0:
		dv.X = 1
	case 1:
		dv.X = -1
	case 2:
		dv.Y = 1
	case 3:
		dv.Y = -1
	}
	return MakeWalkAction(a, dv)
}
