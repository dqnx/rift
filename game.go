package main

import "image/color"

// Game stores and handles all game data.
type Game struct {
	actors       []Actor
	currentActor int

	background color.RGBA
	foreground color.RGBA
}

// Init generates a default actor list.
func (g *Game) Init() {
	for i := 0; i < 5; i++ {
		g.actors = append(g.actors, Actor{ID: i})
	}
}

// Process will update an actor.
func (g *Game) Process() {
	g.actors[g.currentActor].Update()
	g.currentActor = (g.currentActor + 1) % len(g.actors)
}
