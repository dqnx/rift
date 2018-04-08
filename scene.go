package main

import (
	"gitlab.com/rauko1753/gorl"
)

// Scene represents a state of the game. Each scene has its own logic and rendering.
type Scene interface {
	Init()
	HandleInput(gorl.Key) (Transition, Scene)
	Render()
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
}

// Init adds a scene to initialize the scene list.
func (m *SceneManager) Init(s Scene) {
	m.scenes = append(m.scenes, s)
	m.scenes[0].Init()
}

// Empty returns if the scene list is empty.
func (m *SceneManager) Empty() bool {
	return len(m.scenes) == 0
}

// HandleKeyEvent maps the transitions to scene slice manipulations.
func (m *SceneManager) HandleKeyEvent(k gorl.Key) {
	if m.Empty() {
		panic("HandleKeyEvent sees an empty scene list!")
	}

	t, s := m.scenes[0].HandleInput(k)

	switch t {
	case Stay:
	case Next:
		if len(m.scenes) > 1 {
			_, m.scenes = m.scenes[0], m.scenes[1:]
			m.scenes[0].Init()
		} else {
			// Removing one element makes it empty anyways.
			m.scenes = nil
		}
	case Append:
		m.scenes = append(m.scenes, s)
		_, m.scenes = m.scenes[0], m.scenes[1:]
		m.scenes[0].Init()
	case Insert:
		insert := []Scene{s}
		m.scenes = append(insert, m.scenes...)
		m.scenes[0].Init()
	}

}

// RenderAll calls render on each scene
func (m *SceneManager) RenderAll() {
	for _, s := range m.scenes {
		s.Render()
	}
}
