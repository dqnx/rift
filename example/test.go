package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Game logic and pixel communcation channels
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	inputBuffer := make(chan rune)
	renderBuffer := make(chan []int)
	const (
		sleep    = 2
		maxSleep = 10
	)
	totalSleep := 0
	// Logic loop
	go func() {
		fmt.Println("Updating game logic")
		for {
			input := <-inputBuffer
			r := make([]int, 3)
			r[0] = int(input)
			if input == '1' {
				return
			} else if input != '0' {
				fmt.Println("Received input:", input)
				//handle input
				//update logic
			}
			renderBuffer <- r
		}
	}()

	// Pixel: input and rendering loop
	pixel := func() {
		for {
			if totalSleep > maxSleep {
				inputBuffer <- '1'
				return
			} else if r1.Intn(2) > 0 {
				inputBuffer <- 'A'
			} else {
				inputBuffer <- '0'
			}
			r := <-renderBuffer
			//render(r)
			fmt.Println("Received render info:", r)
			fmt.Println("Sleep:", totalSleep)
			time.Sleep(sleep * time.Second)
			totalSleep += sleep
		}
	}
	wrapper(pixel)
}

func wrapper(run func()) {
	run()
}
