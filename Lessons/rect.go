package main

import sf "github.com/zyedidia/sfml/v2.3/sfml"
import "runtime"

const (
	screenWidth  = 800
	screenHeight = 600
)

var textures map[string]*sf.Texture

func LoadTexture(filename string) {
	texture := sf.NewTexture (filename)
	textures[filename] = texture
}

func main() {
	runtime.LockOSThread()

	textures = make (map[string]*sf.Texture)

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	LoadTexture ("playerShip2_red.png")
	sprite := sf.NewSprite (textures["playerShip2_red.png"])

	size := sprite.GetGlobalBounds ();
	sprite.SetOrigin (sf.Vector2f{size.Width/2, size.Height/2})

	sprite.SetPosition (sf.Vector2f {400, 300})

	for window.IsOpen() {
		if event := window.PollEvent(); event != nil {
            switch event.Type {
            case sf.EventClosed:
                window.Close()
            }
        }

        sprite.Rotate (5)

		window.Clear(sf.ColorWhite)
		window.Draw (sprite)
		window.Display()
	}
}