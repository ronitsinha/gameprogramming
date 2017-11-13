package main

import (
	"runtime"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"time"
	"math/rand"
)

func main() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UTC().UnixNano())

	res = NewResources ()

	playerSpriteIndex = 0;

	LoadMap ()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Pacman", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(50)

	player = NewPlayer (sf.Vector2f{screenWidth/2, screenHeight/1.5 - 2})
	blinky := NewEnemy ("blinky", sf.Vector2f{screenWidth/2, screenHeight/2 - 10}, 20)
	pinky := NewEnemy ("pinky", sf.Vector2f{screenWidth/2, screenHeight/2 - 10}, 20)
	inky := NewEnemy ("inky", sf.Vector2f{screenWidth/2, screenHeight/2 - 10}, 20)
	clyde := NewEnemy ("clyde", sf.Vector2f{screenWidth/2, screenHeight/2 - 10}, 22)

	ghosts = append (ghosts, blinky, pinky, inky, clyde)


	var deltaTime float32

	go NextGhost ()

	for window.IsOpen() {
		startTime := time.Now ()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		if enemySpriteIndex >= 30 {
			enemySpriteIndex = 0
		} else {
			enemySpriteIndex ++
		}

		player.Update ()

		for _, g := range ghosts {
			g.Update (deltaTime)
		}

		window.Clear(sf.ColorBlack)

		/*for x, _ := range grid {
			for y,_ := range grid[x] {
				window.Draw (grid[x][y].RectangleShape)
			}
		}*/

		window.Draw (sf.NewSprite (sf.NewTexture("./assets/images/maze.png")))
		
		for _, p := range pacdots {
			window.Draw (p)
		}

		for _, g := range ghosts {
			window.Draw (g)
		}

		window.Draw (player)

		window.Display()

		elapsedTime := time.Since (startTime)
		deltaTime = float32 (elapsedTime) / float32 (time.Second)
	}
}
