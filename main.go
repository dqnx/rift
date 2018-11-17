package main

import (
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/bmp"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := bmp.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Rift",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := loadPicture("tileset.bmp")
	if err != nil {
		panic(err)
	}

	const (
		spritesheetRows = 16
		spritesheetCols = 16
		tileX           = 16
		tileY           = 16
	)

	var tileSprites [spritesheetCols * spritesheetRows]*pixel.Sprite

	for x, i := spritesheet.Bounds().Min.X, 0; x < spritesheet.Bounds().Max.X; x += tileX {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += tileY {
			tileSprites[i] = pixel.NewSprite(spritesheet, pixel.R(x, y, x+tileX, y+tileY))
			i++
		}
	}

	// Tile is a "char" enum.
	type Tile int
	const (
		At Tile = iota
		Hash
		Dot
	)

	type Entity struct {
		Tile
		Name string
	}

	//last := time.Now()
	for !win.Closed() {
		//dt := time.Since(last).Seconds()
		//last = time.Now()

		win.Clear(colornames.Forestgreen)
	}
}

func main() {
	pixelgl.Run(run)
}
