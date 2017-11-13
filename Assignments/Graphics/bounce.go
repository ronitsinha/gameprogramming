package main

import sf "github.com/manyminds/gosfml"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Bouncing Square", sf.StyleDefault, sf.DefaultContextSettings())
	window.SetVSyncEnabled(true)
	window.SetFramerateLimit(60)

	square, _ := sf.NewRectangleShape ()
	square.SetSize (sf.Vector2f {30, 30})
	square.SetFillColor(sf.ColorBlue())
	square.SetOrigin (sf.Vector2f {15, 15})
	square.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	v := sf.Vector2f {5, -5}

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
			v.X = -v.X
		} else if square.GetPosition().X <= 0 {
			v.X = v.X
		}

		if square.GetPosition ().Y >= screenHeight {
			v.Y = -5
		} else if square.GetPosition ().Y <= 0 {
			v.Y = 5
		}

		square.Move (v)
	}
}