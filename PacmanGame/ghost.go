package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"strconv"
)

type Enemy struct {
	*sf.Sprite

	path []*Node
	moving bool
	targetIndex int
	name string
	sprDirection int
	isActive bool
	speed float32
}

func NewEnemy (name string, pos sf.Vector2f, spd float32) *Enemy {
	enmy := new (Enemy)

	enmy.path = make ([]*Node, 0)
	enmy.sprDirection = 0
	enmy.isActive = false

	enmy.Sprite = sf.NewSprite (res.images[name + strconv.Itoa (enmy.sprDirection) + "_0.png"])
	enmy.name = name
	enmy.SetOrigin (sf.Vector2f {enmy.GetGlobalBounds ().Width/2, enmy.GetGlobalBounds ().Height/2})

	enmy.SetPosition (pos)
	enmy.speed = spd

	return enmy
}

func (e *Enemy) Update (dt float32) {	
	if e.isActive {
		for x, _ := range grid {
			for _, n := range grid[x] {
				if Contains (e.path, n) {
					n.SetFillColor (sf.ColorGreen)
				} else {
					n.SetFillColor (sf.ColorBlack)
				}
			}
		}

		e.path = FindPath (e.GetPosition (), player.GetPosition ())
		go e.MoveAlongPath (dt)
	}
	e.Animate ()
}

func (e *Enemy) Animate() {
	var frame int

	if enemySpriteIndex > 15 {
		frame = 1
	} else {
		frame = 0
	}

	e.SetTexture (res.images[e.name + strconv.Itoa (e.sprDirection) + "_" + strconv.Itoa (frame) + ".png"], false)
}

func (e *Enemy) MoveAlongPath(dt float32) {
	atPlayer, _ := e.GetGlobalBounds ().Intersects (player.GetGlobalBounds ())

	if atPlayer || len (e.path) <= 0 {
		return
	}

	currentWaypoint := e.path[0].GetPosition ()

	for true {
		if len (e.path) <= 0 {
			return
		}

		if e.GetPosition ().IsEqual (currentWaypoint) {
			e.targetIndex ++
			if e.targetIndex >= len (e.path) {
				return
			}

			currentWaypoint = e.path [e.targetIndex].GetPosition ()
		}

		vx := (currentWaypoint.X + NODE_SIZE/2) - e.GetPosition ().X
		vy := (currentWaypoint.Y + NODE_SIZE/2) - e.GetPosition ().Y

		if int (vx) > 0 {
			vx = 1
			e.sprDirection = 0
		} else if int (vx) < 0 {
			vx = -1
			e.sprDirection = 3
		} else {
			vx = 0
		}

		if int (vy) > 0 {
			vy = 1
			e.sprDirection = 2
		} else if int (vy) < 0 {
			vy = -1
			e.sprDirection = 1
		} else {
			vy = 0
		}

		for _, g := range ghosts {
			if e != g {
				if g.GetGlobalBounds ().Contains (sf.Vector2f {e.GetPosition ().X + vx, e.GetPosition ().Y + vy}) && g.isActive {
					return
				}
			}
		}

		e.Move(sf.Vector2f {vx * dt * e.speed, vy * dt * e.speed})

		e.targetIndex = 0

		return
	}
}
