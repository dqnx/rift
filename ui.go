package main

import (
	"fmt"
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// Tile is a "char" enum.
type Tile int

const (
	At   = Tile(176)
	Hash = Tile(211)
	Dot  = Tile(9)
)

type Entity struct {
	Tile
	Name string
}

type Position struct {
	X int
	Y int
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// LaunchUI creates the game window and starts the input handling.
func LaunchUI() {
	cfg := pixelgl.WindowConfig{
		Title:  "Rift",
		Bounds: pixel.R(0, 0, 256, 256),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := loadPicture("Zilk_16x16.png")
	if err != nil {
		panic(err)
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	const (
		spritesheetRows = 16
		spritesheetCols = 16
		tileX           = 16
		tileY           = 16
		tileXh          = tileX / 2
		tileYh          = tileY / 2
	)

	var tileSprites [spritesheetCols * spritesheetRows]*pixel.Sprite
	for y, i := spritesheet.Bounds().Min.Y, 0; y < spritesheet.Bounds().Max.Y; y += tileY {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += tileX {
			tileSprites[i] = pixel.NewSprite(spritesheet, pixel.R(x, y, x+tileX, y+tileY))
			i++
		}
	}

	tiles := make(map[Position]Tile)
	const (
		testMapX = 4
		testMapY = 6
	)
	inputMap := func() [testMapX * testMapY]int {
		return [...]int{
			211, 211, 211, 211,
			211, 9, 9, 211,
			211, 9, 9, 211,
			211, 9, 9, 211,
			211, 9, 9, 211,
			211, 211, 211, 211}
	}
	for i, t := range inputMap() {
		p := Position{X: i % testMapX, Y: i / testMapX}
		fmt.Println(p)
		tiles[p] = Tile(t)
	}

	//last := time.Now()

	for !win.Closed() {
		batch.Clear()
		for pos, tile := range tiles {
			spritePos := pixel.IM.Moved(pixel.V(float64(pos.X*tileX+tileXh), float64(pos.Y*tileY+tileYh)))
			tileSprites[tile].Draw(batch, spritePos)
		}
		//dt := time.Since(last).Seconds()
		//last = time.Now()
		win.Clear(colornames.Black)
		batch.Draw(win)
		win.Update()
	}
}
