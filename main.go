package main

import (
	"fmt"
	"strconv"
	"time"

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
	scenes.Init(TitleScene{})

	// Handle Input
	// Waits for input (key press)
	fmt.Println("Starting Loop")
	go func() {
		for running {
			if scenes.Empty() {
				fmt.Println("Scenes Empty, exiting input loop.")
				running = false
			} else {
				scenes.HandleKeyEvent(gorl.TermGetKey())
			}
		}
		fmt.Println("Input loop done.")
	}()

	// Rendering loop
	tick := time.Tick(frameRate * time.Microsecond)
	frameStart := time.Now()

	for running {
		// Measure the actual fps
		now := time.Now()
		frameTime := now.Sub(frameStart)
		frameStart = now
		fps := 1.0 / frameTime.Seconds()

		<-tick
		// Clear
		gorl.TermClear()

		// Draw
		// Row 0 is reserved for status
		TermDrawText(0, 0, strconv.FormatFloat(fps, 'f', 1, 64))
		scenes.RenderAll()

		// Refresh
		gorl.TermRefresh()

	}
	fmt.Println("Render loop done.")
}
