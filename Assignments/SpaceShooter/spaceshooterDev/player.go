package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math"
	"time"
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
	*sf.Sprite

	speed float32
	// An array with a fixed length of 5
	keys [5]sf.KeyCode

	canShoot bool
}

func NewPlayer(keys [5]sf.KeyCode, pos sf.Vector2f, name string) *Player {
	player := new (Player)
	player.Sprite = sf.NewSprite (res.images[name + ".png"])
	player.SetOrigin (sf.Vector2f {player.GetGlobalBounds ().Width/2, player.GetGlobalBounds ().Height/2})

	player.SetPosition (pos)

	player.keys = keys
	player.canShoot = true

	return player
}

func (p *Player) Shoot() {
	lasers = append (lasers, NewLaser (p.GetPosition (), p.GetRotation (), LASERSPEED), NewLaser (p.GetPosition (), p.GetRotation ()-15, LASERSPEED), NewLaser (p.GetPosition (), p.GetRotation ()+15, LASERSPEED))

	soundBuffer := res.sounds ["sfx_laser1.ogg"]
	sound := sf.NewSound (soundBuffer)

	sound.Play ()

	go func() {
		time.Sleep (SHOOTCOOLDOWN*time.Millisecond)
		p.canShoot = true
	}()
}

func (p *Player) Update (dt float32) {
	if sf.KeyboardIsKeyPressed (p.keys[0]) {
		if p.speed < SHIPMAXSPEED {
			p.speed += SHIPACCELERATION
		}
	} else if p.speed > 0 {
		p.speed -= SHIPDECCELERATION
	}
	if sf.KeyboardIsKeyPressed (p.keys[2]) {
		p.Rotate (-SHIPROTATESPEED * dt * 60)
	}
	if sf.KeyboardIsKeyPressed (p.keys[3]) {
		p.Rotate (SHIPROTATESPEED * dt * 60)
	}

	if sf.KeyboardIsKeyPressed (p.keys[4]) && p.canShoot {
		p.Shoot ()

		p.canShoot = false
	}



	if p.speed < 0 {
		p.speed = 0
	} else if p.speed > SHIPMAXSPEED {
		p.speed = SHIPMAXSPEED
	}

	angle := p.GetRotation () - 90
	angleRad := angle * math.Pi/180
		
	vx := p.speed * float32 (math.Cos (float64 (angleRad)))
	vy := p.speed * float32 (math.Sin (float64(angleRad)))
	p.Move (sf.Vector2f {vx * dt * 60, vy * dt * 60})

	Wrap (p.Sprite)
}
