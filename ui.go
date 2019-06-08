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
	const (
		spritesheetRows      = 16
		spritesheetCols      = 16
		spriteNum       byte = spritesheetCols*spritesheetRows - 1
		tileX                = 16
		tileY                = 16
		tileXh               = tileX / 2
		tileYh               = tileY / 2
		screenX              = 512
		screenY              = 256
	)

	cfg := pixelgl.WindowConfig{
		Title:  "Rift",
		Bounds: pixel.R(0, 0, screenX, screenY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := loadPicture("assets/tileset.png")
	if err != nil {
		panic(err)
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	var tileSprites [spritesheetCols * spritesheetRows]*pixel.Sprite
	for y, i := spritesheet.Bounds().Max.Y-tileY, 0; y > spritesheet.Bounds().Min.Y-tileY; y -= tileY {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += tileX {
			tileSprites[i] = pixel.NewSprite(spritesheet, pixel.R(x, y, x+tileX, y+tileY))
			i++
		}
	}

	size := dim.V(screenX/tileX, screenY/tileY)
	scenes := SceneManager{size: size}
	scenes.Init(&TitleScene{})

	game := &Game{background: colornames.Darkgrey, foreground: colornames.White, size: size.Add(dim.V(-2, -2))}
	game.Init()
	for !win.Closed() {
		scenes.HandleKeyEvent(win)
		if scenes.Empty() {
			break
		}

		err := scenes.Update(game)
		if err != nil {
			panic(err)
		}
		sceneTiles, err := scenes.Render()
		if err != nil {
			panic(err)
		}

		batch.Clear()
		for _, tiles := range sceneTiles {
			for pos, tile := range tiles {
				spritePos := pixel.IM.Moved(pixel.V(float64(pos.X*tileX+tileXh), float64(pos.Y*tileY+tileYh)))
				tileSprites[tile].Draw(batch, spritePos)
			}
		}

		win.Clear(colornames.Black)
		batch.Draw(win)
		win.Update()
	}
}
