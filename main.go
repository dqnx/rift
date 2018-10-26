package main

func main() {
	// Launch the game loop
	update := make(chan bool)
	gameRunning := true
	go func() {
		for gameRunning {
			<-update
		}
	}()

	// Create the window

}
