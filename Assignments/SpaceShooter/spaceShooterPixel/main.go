package main

import (
	pixel "github.com/faiface/pixel"
	pixelgl "github.com/faiface/pixel/pixelgl"
	color "golang.org/x/image/colornames"

	"time"
)

const (
	screenWidth = 1024
	screenHeight = 768
)

var res *Resources

func run() {
	cfg := pixelgl.WindowConfig {
		Title: "Test",
		Bounds: pixel.R (0, 0, screenWidth, screenHeight),
		VSync: true,
	}

	window, err := pixelgl.NewWindow (cfg)
	if err != nil {
		panic (err)
	}

	// Smoothing of rotated images
	window.SetSmooth (true)

	player := NewPlayer ([4]pixelgl.Button {pixelgl.KeyUp, pixelgl.KeyDown, pixelgl.KeyLeft, pixelgl.KeyRight}, pixel.IM.Moved(window.Bounds ().Center()), "playerShip1_blue.png")

	last := time.Now ()
	for !window.Closed () {
		dt := time.Since (last).Seconds ()
		last = time.Now ()

		player.Update (float32 (dt), window)

		window.Clear (color.Black)

		player.Draw(window)

		window.Update ()
	}
}

func main() {
	res = NewResources ()
	pixelgl.Run (run)
}