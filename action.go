package main

// Action is a command pattern and returns a callback for something to do with an Actor.
type Action interface {
	Perform(*Actor)
}

func WalkAction(a *Actor, x int, y int) func() {
	return func() {
		a.Move(x, y)
	}
}
