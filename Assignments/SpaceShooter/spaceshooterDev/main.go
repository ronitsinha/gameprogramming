package main

import (
	"runtime"
	"time"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math/rand"
)

var res *Resources
var lasers []*Laser
var asteroids []*Asteroid
var explosions []*Explosion

const (
	screenWidth  = 800
	screenHeight = 600
)

func random(min, max float32) float32 {
    return rand.Float32()*(max-min) + min
}

func main() {
	runtime.LockOSThread()
	rand.Seed(time.Now().UTC().UnixNano())

	res = NewResources ()
	
	asteroids = make ([]*Asteroid, 0)
	lasers = make ([]*Laser, 0)

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Space Shooter", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	player := NewPlayer ([5]sf.KeyCode {sf.KeyUp, sf.KeyDown, sf.KeyLeft, sf.KeyRight, sf.KeySpace}, sf.Vector2f{screenWidth/2, screenHeight/2}, "playerShip1_blue")
	player1 := NewPlayer ([5]sf.KeyCode {sf.KeyW, sf.KeyS, sf.KeyA, sf.KeyD, sf.KeyLShift}, sf.Vector2f{screenWidth/2, screenHeight/2}, "playerShip3_red")


	for i := 0; i < NUMASTEROIDS; i ++ {
		appendAstr := NewAsteroid (sf.Vector2f {random (0, screenWidth), random (0, screenHeight)}, random (100, 300), random (0, 360), random (3, 7))
		asteroids = append (asteroids, appendAstr)
	}


	var dt float32

	for window.IsOpen() {
		start := time.Now ()

		if event := window.PollEvent(); event != nil {
			switch event.Type {
			case sf.EventClosed:
				window.Close()
			}
		}

		player.Update (dt)
		player1.Update (dt)

		for _, l := range lasers {
			l.Update (dt)
		}

		for _, a := range asteroids {
			a.Update (dt)
		}

		for _, l := range lasers {
			for _, a := range asteroids {
				if Intersects (a.Sprite, l.Sprite) {
					l.dead = true
					a.dead = true

					SpawnExplosion (a.GetPosition ())

					soundBuffer := res.sounds ["sfx_explosion.wav"]
					sound := sf.NewSound (soundBuffer)
					sound.Play ()
				}
			}	
		}

		window.Clear(sf.ColorBlack)

		for _, e := range explosions {
			window.Draw (e)
		}

		for _, l := range lasers {
			window.Draw (l)
		}

		for _, a := range asteroids {
			window.Draw (a)
		}

		window.Draw (player)
		window.Draw (player1)

		var tempLasers []*Laser

		for _, l := range lasers {
			if !l.dead {
				tempLasers = append (tempLasers, l)
			}
		}

		lasers = tempLasers

		var tempAst []*Asteroid

		for _, a := range asteroids {
			if !a.dead {
				tempAst = append (tempAst, a)
			}
		}

		asteroids = tempAst

		var tempExp []*Explosion

		for _, e := range explosions {
			if !e.dead {
				tempExp = append (tempExp, e)
			}
		}

		explosions = tempExp

		window.Display()
		elapsed := time.Since (start)
		dt = float32 (elapsed) / float32 (time.Second)
	}
}
