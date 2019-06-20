package main

import (
	"container/heap"
	"fmt"
	"image/color"
	"math/rand"

	dim "github.com/sean-brock/dimension"
)

// Game stores and handles all game data.
type Game struct {
	actors       map[ID]*Actor
	energy       PriorityQueue
	turnEnergy   int
	currentActor ID

	background color.RGBA
	foreground color.RGBA

	size dim.Vec
}

func (g *Game) Size() dim.Vec {
	return g.size
}

func (g *Game) Actors() map[ID]*Actor {
	return g.actors
}

// Init generates a default actor list.
func (g *Game) Init() {
	// Setting 1 energy but not initializing it!
	g.energy = make(PriorityQueue, 1)
	heap.Init(&g.energy)

	g.actors = make(map[ID]*Actor)
	for i := 0; i < 5; i++ {
		fmt.Println("actor:", i)
		a := new(Actor)
		a.ID = ID(i)
		a.Position.X = rand.Intn(g.size.X)
		a.Position.Y = rand.Intn(g.size.Y)
		a.Speed = rand.Intn(5) + 1
		g.actors[ID(i)] = a

		addActor(&g.energy, a)
	}

	g.turnEnergy = 10
	AddPriority(g.energy, g.turnEnergy)
}

// addActor pushes an actor to a queue.
func addActor(pq *PriorityQueue, a *Actor) {
	item := &Item{
		ID:       a.ID,
		priority: 0,
	}
	heap.Push(pq, item)
}

// Process will update an actor.
func (g *Game) Process() {
	energized := heap.Pop(&g.energy).(*Item)
	// Validate ID exists
	actor := g.actors[energized.ID]
	// Check if priority is above energy for a turn.
	if energized.Priority() > g.turnEnergy {
		actor.Update()
		action := actor.Action()
		success := action.Perform(g)
		// If failed the action, add back to queue
		if success {
			addActor(&g.energy, actor)
			AddPriority(g.energy, g.turnEnergy)
		} else {
			heap.Push(&g.energy, energized)
		}
	} else {
		// If no actor has enough energy, add back the actor and
		// add the turn energy.
		heap.Push(&g.energy, energized)
		AddPriority(g.energy, g.turnEnergy)
	}
}
