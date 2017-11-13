package main

import (
	sf "github.com/manyminds/gosfml"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Gravity", sf.StyleDefault, sf.DefaultContextSettings())
	window.SetVSyncEnabled(true)
	window.SetFramerateLimit(60)

	square, _ := sf.NewRectangleShape ()
	square.SetSize (sf.Vector2f {30, 30})
	square.SetFillColor(sf.ColorBlue())
	square.SetOrigin (sf.Vector2f {15, 15})
	square.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	grav := float32 (1)
	v := sf.Vector2f {5, 10}

	for window.IsOpen() {
		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					window.Close()
				}
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.ColorWhite())
		window.Draw(square, sf.DefaultRenderStates())
		window.Display()
	
		if square.GetPosition().X >= screenWidth {
			v.X = -5
		} else if square.GetPosition().X <= 0 {
			v.X = 5
		}

		if square.GetPosition ().Y >= screenHeight {
			// Hits the bottom
			v.Y = -20
		} else {
			// On the way down
			v.Y += grav
		}

		square.Move (v)
	}
}