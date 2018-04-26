package main

import (
	"fmt"
	"gitlab.com/rauko1753/gorl"
)

// Need to wrap in mainthread...
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
	fmt.Println("Starting Loop")
	for running {
		if scenes.Empty() {
			fmt.Println("Scenes Empty, exiting loop.")
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
	fmt.Println("Render loop done.")
}
