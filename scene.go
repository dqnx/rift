package main

import (
	"errors"

	"github.com/faiface/pixel/pixelgl"
	dim "github.com/sean-brock/dimension"
)

// Scene represents a state of the game. Each scene has its own logic and rendering.
type Scene interface {
	Init(size dim.Vec)
	HandleInput(w *pixelgl.Window) (Transition, Scene)
	ExportTiles() TileList
	ExportText() TextList
}

// Transition is a message, describing what to do after a scene frame.
type Transition int

const (
	// Stay does nothing
	Stay Transition = iota
	// Next removes curent scene and goes to the top scene, if it exists.
	Next
	// Append adds a scene at the end and pops the current scene.
	Append
	// Insert a scene at the beginning and switch to it.
	Insert
)

// SceneManager allows scenes to be swapped and added.
type SceneManager struct {
	scenes []Scene
	size   dim.Vec
}

// Init adds a scene to initialize the scene list.
func (m *SceneManager) Init(s Scene) {
	m.scenes = append(m.scenes, s)
	m.scenes[0].Init(m.size)
}

// Empty returns if the scene list is empty.
func (m *SceneManager) Empty() bool {
	return len(m.scenes) == 0
}

// HandleKeyEvent maps the transitions to scene slice manipulations.
func (m *SceneManager) HandleKeyEvent(w *pixelgl.Window) {
	if m.Empty() {
		panic("HandleKeyEvent sees an empty scene list!")
	}

	t, s := m.scenes[0].HandleInput(w)

	switch t {
	case Stay:
	case Next:
		if len(m.scenes) > 1 {
			_, m.scenes = m.scenes[0], m.scenes[1:]
			m.scenes[0].Init(m.size)
		} else {
			// Removing one element makes it empty anyways.
			m.scenes = nil
		}
	case Append:
		m.scenes = append(m.scenes, s)
		_, m.scenes = m.scenes[0], m.scenes[1:]
		m.scenes[0].Init(m.size)
	case Insert:
		insert := []Scene{s}
		m.scenes = append(insert, m.scenes...)
		m.scenes[0].Init(m.size)
	}

}

// RenderAll calls render on each scene
func (m *SceneManager) Render() ([]TileList, error) {
	if len(m.scenes) > 0 {
		tiles := make([]TileList, len(m.scenes))
		for i, s := range m.scenes {
			tiles[i] = s.ExportTiles()
		}
		return tiles, nil
	}
	return nil, errors.New("failed to render no scenes")
}
