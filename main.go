package main

import (
	"time"
)

// Need to wrap in mainthread...
func main() {
	const frameRate = 60 * time.Millisecond
	
	// Handle Input
	// Waits for input (key press)
	/*
	go func() {
		for { // window not closed
			handleInput(getEvent())
		}
	}()
	*/
	
	// Rendering loop
	tick := time.Tick(frameRate)
	for { // window not closed
		<-tick
		// Clear
		// Draw
		// Refresh
	}	
}
