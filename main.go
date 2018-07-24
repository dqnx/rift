package main

import (
	"gitlab.com/rauko1753/gorl"
)

//const debug = true
const debug = false

func main() {

	// Frametime in microseconds
	const frameRate = 1000000 / 60

	gorl.TermMustInit()
	defer gorl.TermDone()
	running := true

	var scenes SceneManager
	scenes.Init(&TitleScene{})

	// Handle Input
	// Waits for input (key press)
	for running {
		if scenes.Empty() {
			running = false
			break
		}

		gorl.TermClear()

		// Draw
		// Row 0 is reserved for status
		//TermDrawText(0, 0, strconv.FormatFloat(fps, 'f', 1, 64))
		scenes.RenderAll()

		// Refresh
		gorl.TermRefresh()

		// Place at end, so first loop will draw immediately.
		scenes.HandleKeyEvent(gorl.TermGetKey())
	}
}
