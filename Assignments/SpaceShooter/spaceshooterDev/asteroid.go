package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math"
	"math/rand"
	"strconv"
)

const (
	NUMASTEROIDS = 5
)

type Asteroid struct {
	*sf.Sprite

	speed float32
	moveRotation float32
	rotationSpeed float32
	dead bool
}

func NewAsteroid (pos sf.Vector2f, rotationSpeed, moveAngle, speed float32) *Asteroid {
	astrd := new(Asteroid)

	index := rand.Intn (4) + 1
	astrd.Sprite = sf.NewSprite (res.images["asteroid" + strconv.Itoa (index) + ".png"])
	astrd.SetOrigin (sf.Vector2f {astrd.GetGlobalBounds ().Width/2, astrd.GetGlobalBounds ().Height/2})

	astrd.SetPosition (pos)
	astrd.moveRotation = moveAngle
	astrd.rotationSpeed = rotationSpeed
	astrd.speed = speed

	return astrd
}

func (a *Asteroid) Update (dt float32) {
	a.Rotate (a.rotationSpeed * dt)

	angle := a.moveRotation - 90
	angleRad := angle * math.Pi/180
		
	vx := a.speed * float32 (math.Cos (float64 (angleRad)))
	vy := a.speed * float32 (math.Sin (float64(angleRad)))
	a.Move (sf.Vector2f {vx * dt * 60, vy * dt * 60})	

	Wrap (a.Sprite)
}