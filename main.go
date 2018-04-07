package main

import (
	"fmt"
	"time"

	"gitlab.com/rauko1753/gorl"
)

// Need to wrap in mainthread...
func main() {
	const frameRate = 60 * time.Millisecond

	gorl.TermMustInit()
	defer gorl.TermDone()
	running := true

	// Handle Input
	// Waits for input (key press)
	go func() {
		for running {
			switch gorl.TermGetKey() {
			case gorl.KeyEsc:
				running = false
			}
		}
		fmt.Println("Input loop done.")
	}()

	// Rendering loop
	tick := time.Tick(frameRate)
	for running {
		<-tick
		// Clear
		gorl.TermClear()
		// Draw
		// Refresh
		gorl.TermRefresh()
	}
	fmt.Println("Render loop done.")
}
