package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math"
	//"fmt"
)

const (
	LASERSPEED = 10
)

type Laser struct {
	*sf.Sprite 

	speed float32
	dead bool
}

func NewLaser (pos sf.Vector2f, rotation float32, speed float32) *Laser {
	lser := new (Laser)

	lser.Sprite = sf.NewSprite (res.images["laserBlue.png"])
	lser.SetOrigin (sf.Vector2f {lser.GetGlobalBounds ().Width/2, lser.GetGlobalBounds ().Height/2})
	lser.SetPosition (pos)

	lser.SetRotation (rotation)
	lser.speed = speed

	return lser
}

func (ls *Laser) Update(dt float32) {
	x := ls.GetPosition ().X
	y := ls.GetPosition ().Y

	size := ls.GetGlobalBounds ()

	if x < -size.Width/2 || x > screenWidth + size.Width/2 || y < -size.Height/2 || y > screenHeight + size.Height/2 {
		ls.dead = true
	}

	angle := ls.GetRotation () - 90
	angleRad := angle * math.Pi/180
		
	vx := ls.speed * float32 (math.Cos (float64 (angleRad)))
	vy := ls.speed * float32 (math.Sin (float64(angleRad)))
	ls.Move (sf.Vector2f {vx * dt * 60, vy * dt * 60})	
}
