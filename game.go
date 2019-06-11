package main

import (
	"container/heap"
	"image/color"
	"math/rand"

	dim "github.com/sean-brock/dimension"
)

// Game stores and handles all game data.
type Game struct {
	actors       map[ID]*Actor
	energy       PriorityQueue
	currentActor ID

	background color.RGBA
	foreground color.RGBA

	size dim.Vec
}

func (g *Game) Size() dim.Vec {
	return g.size
}

func (g *Game) Actors() map[ID]Actor {
	return g.actors
}

// Init generates a default actor list.
func (g *Game) Init() {
	for i := 0; i < 5; i++ {
		a := Actor{ID: i}
		a.Position.X = rand.Intn(g.size.X)
		a.Position.Y = rand.Intn(g.size.Y)
		g.actors[ID(i)] = a
	}

	g.energy = make(PriorityQueue, 1)
	heap.Init(&g.energy)
}

// Process will update an actor.
func (g *Game) Process() {
	energized := heap.Pop(&g.energy).(*Item)
	// Validate ID exists
	actor := g.actors[energized.ID]
	actor.Update()
	action := actor.Action()
	success := action.Perform(g)
	if !success {
		heap.Push(&g.energy, energized)
	}
}
