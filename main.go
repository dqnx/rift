package main

import (
	"fmt"

	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/consolecolor"
	"github.com/BigJk/ramen/font"
	"github.com/BigJk/ramen/t"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func main() {
	// load a font you like
	font, err := font.New("fonts/terminus-11x11.png", 11, 11)
	if err != nil {
		panic(err)
	}

	// create a 50x30 cells console with the title 'ramen example'
	con, err := console.New(50, 30, font, "ramen example")
	if err != nil {
		panic(err)
	}

	game := &Game{}
	game.Init()
	// set a tick hook. This function will be executed
	// each tick (60 ticks per second by default) even
	// when the fps is lower than 60fps. This is a good
	// place for your game logic.
	//
	// The timeDelta parameter is the elapsed time in seconds
	// since the last tick.

	con.SetTickHook(func(timeElapsed float64) error {
		// your game logic
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			command := WalkAction()
		}

		return nil
	})

	// set a pre-render hook. This function will be executed
	// each frame before the drawing happens. This is a good
	// place to draw onto the console, because it only executes
	// if a draw is really about to happen.
	//
	// The timeDelta parameter is the elapsed time in seconds
	// since the last frame.
	con.SetPreRenderHook(func(screen *ebiten.Image, timeDelta float64) error {
		con.ClearAll(t.Background(consolecolor.New(50, 50, 50)))
		con.Print(2, 2, "Hello!\nTEST\n Line 3", t.Foreground(consolecolor.New(0, 255, 0)), t.Background(consolecolor.New(255, 0, 0)))
		con.Print(2, 7, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nElapsed: %0.4f", ebiten.CurrentFPS(), ebiten.CurrentFPS(), timeDelta))
		return nil
	})

	// start the console with a scaling of 1
	con.Start(1)
}

func handleInput() {

}
