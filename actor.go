package main

import "fmt"

// Actor is any game character or sentient creature.
type Actor struct {
	Name string
	ID   int
	X    int
	Y    int
}

// Update will return an action from an actor.
func (a *Actor) Update() {
	fmt.Println("Updated Actor:", a.ID)
}

// Move changes an Actor's position.
func (a *Actor) Move(x, y int) {
	a.X = x
	a.Y = y
}
