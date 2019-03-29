package main

import (
	"image"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	dim "github.com/sean-brock/dimension"
	"golang.org/x/image/colornames"
)

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

	spritesheet, err := loadPicture("tileset.png")
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

	screen := MakeScreen(dim.V(6, 6), dim.V(0, 0))
	screen.Set(dim.V(0, 1), W2v)
	screen.Set(dim.V(0, 2), W2v)
	screen.Set(dim.V(0, 3), W2v)
	screen.Set(dim.V(0, 4), W2v)
	screen.Set(dim.V(0, 0), W2v)
	screen.Set(dim.V(0, 5), W2v)
	screen.Set(dim.V(1, 2), At)

	for !win.Closed() {
		tiles := screen.Tiles(dim.V(0, 0))
		batch.Clear()
		for pos, tile := range tiles {
			spritePos := pixel.IM.Moved(pixel.V(float64(pos.X*tileX+tileXh), float64(pos.Y*tileY+tileYh)))
			tileSprites[tile].Draw(batch, spritePos)
		}
		win.Clear(colornames.Black)
		batch.Draw(win)
		win.Update()
	}
}
