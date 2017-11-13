package main

import sf "github.com/manyminds/gosfml"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, sf.DefaultContextSettings())
	window.SetVSyncEnabled(true)
	window.SetFramerateLimit(60)

	squareSmall, _ := sf.NewRectangleShape ()
	squareMed, _ := sf.NewRectangleShape ()  
	squareBig, _ := sf.NewRectangleShape ()
	
	squareSmall.SetSize(sf.Vector2f{30, 30})
	squareMed.SetSize(sf.Vector2f{100, 100})
	squareBig.SetSize(sf.Vector2f{200, 200})

	squareSmall.SetFillColor (sf.ColorWhite ())
	squareMed.SetFillColor (sf.ColorWhite ())
	squareBig.SetFillColor (sf.ColorWhite ())

	squareSmall.SetPosition (sf.Vector2f {screenWidth/2, screenHeight/2})
	squareMed.SetPosition (sf.Vector2f {screenWidth/2, screenHeight/2})
	squareBig.SetPosition (sf.Vector2f {screenWidth/2, screenHeight/2})

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
		
		squareSmall.Rotate (8)
		red := squareSmall.GetFillColor().R
		green := squareSmall.GetFillColor().G
		blue := squareSmall.GetFillColor().B
		plusfactor := uint8(10)
		squareSmall.SetFillColor (sf.Color {red + plusfactor, green + plusfactor *5, blue + plusfactor*10, uint8 (255)})
		
		squareMed.Rotate (5)
		red1 := squareMed.GetFillColor().R
		green1 := squareMed.GetFillColor().G
		blue1 := squareMed.GetFillColor().B
		plusfactor1 := uint8(5)
		squareMed.SetFillColor (sf.Color {red1 + plusfactor1, green1 + plusfactor1*5, blue1 + plusfactor1*10, uint8 (255)})
		
		squareBig.Rotate (1)
		red2 := squareBig.GetFillColor().R
		green2 := squareBig.GetFillColor().G
		blue2 := squareBig.GetFillColor().B
		plusfactor2 := uint8(1)
		squareBig.SetFillColor (sf.Color {red2 + plusfactor2, green2 + plusfactor2*5, blue2 + plusfactor2*10, uint8 (255)})

		window.Clear(sf.ColorWhite())
		window.Draw(squareBig, sf.DefaultRenderStates())
		window.Draw(squareMed, sf.DefaultRenderStates())
		window.Draw(squareSmall, sf.DefaultRenderStates())
		window.Display()
	}
}