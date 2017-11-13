package main 

import (
	pixel "github.com/faiface/pixel"
	pixelgl "github.com/faiface/pixel/pixelgl"
	"math"
	//"time"
	"fmt"
)

const (
	SHIPROTATESPEED = 5
	SHIPMAXSPEED = 10
	SHIPACCELERATION = 0.1
	SHIPDECCELERATION = 0.25
	SHOOTCOOLDOWN = 250
)

type Player struct {
	// embedding a sprite
	*pixel.Sprite

	speed float32
	// An array with a fixed length of 5
	keys [4]pixelgl.Button

	rotation float64

}

func NewPlayer(keys [4]pixelgl.Button, pos pixel.Matrix, name string) *Player {
	p := new (Player)

	p.Sprite = pixel.NewSprite (res.images[name], res.images[name].Bounds ())
	p.keys = keys

	p.SetMatrix (pos)

	return p
}

func (p *Player) Update(dt float32, w *pixelgl.Window) {
	if w.Pressed (p.keys[0]) {
		if p.speed < SHIPMAXSPEED {
			p.speed += SHIPACCELERATION
		}
	} else if p.speed > 0 {
		p.speed -= SHIPDECCELERATION
	}
	if w.Pressed (p.keys[2]) {
		p.SetMatrix (p.Matrix ().Rotated (p.Matrix ().Project (pixel.V (0, 0)), float64 (SHIPROTATESPEED * dt)))
		p.rotation += float64 (SHIPROTATESPEED)
	}
	if w.Pressed (p.keys[3]) {
		p.SetMatrix (p.Matrix ().Rotated (p.Matrix ().Project (pixel.V (0, 0)), float64 (-SHIPROTATESPEED * dt)))
		p.rotation -= float64 (SHIPROTATESPEED)
	}

	if p.rotation > 360 {
		p.rotation = 0
	} else if p.rotation < 0 {
		p.rotation = 360
	}

	/*if sf.KeyboardIsKeyPressed (p.keys[4]) && p.canShoot {
		p.Shoot ()

		p.canShoot = false
	}*/

	if p.speed < 0 {
		p.speed = 0
	} else if p.speed > SHIPMAXSPEED {
		p.speed = SHIPMAXSPEED
	}

	angle := p.rotation//p.Matrix ().Project (pixel.V (10, 10)).Angle ()
	fmt.Println (angle)

	angleRad := angle*math.Pi/180


	vx := p.speed * float32 (math.Cos (float64 (angleRad)))
	vy := p.speed * float32 (math.Sin (float64(angleRad)))

	p.SetMatrix (p.Matrix ().Moved (pixel.V (float64(vx * dt * 60), float64(vy * dt * 60))))

	//Wrap (p.Sprite)
}