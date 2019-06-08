package main

import (
	"image/color"
	"math/rand"

	dim "github.com/sean-brock/dimension"
)

// Game stores and handles all game data.
type Game struct {
	actors       []Actor
	currentActor int

	background color.RGBA
	foreground color.RGBA

	size dim.Vec
}

func (g *Game) Size() dim.Vec {
	return g.size
}

func (g *Game) Actors() []Actor {
	return g.actors
}

// Init generates a default actor list.
func (g *Game) Init() {
	for i := 0; i < 5; i++ {
		a := Actor{ID: i}
		a.Position.X = rand.Intn(g.size.X)
		a.Position.Y = rand.Intn(g.size.Y)
		g.actors = append(g.actors, a)
	}
}

// Process will update an actor.
func (g *Game) Process() {
	g.actors[g.currentActor].Update()
	action := g.actors[g.currentActor].Action()
	success := action.Perform(g)
	if success {
		g.currentActor = (g.currentActor + 1) % len(g.actors)
	}
}
