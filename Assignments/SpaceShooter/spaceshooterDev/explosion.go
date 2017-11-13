package main 

import (
	"strconv"
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"time"
)

type Explosion struct {
	*sf.Sprite

	curFrame int
	dead bool
}

func NewExplosion (pos sf.Vector2f) *Explosion {
	e := new (Explosion)
	e.Sprite = sf.NewSprite (res.images["explosion_0.png"])

	e.SetOrigin (sf.Vector2f {e.GetGlobalBounds ().Width/2, e.GetGlobalBounds ().Height/2})

	e.SetPosition (pos)

	go e.Animate ()

	return e
}

func (e *Explosion) Animate() {
	for i := 0; i < 24; i ++ {
		e.Sprite.SetTexture (res.images["explosion" + strconv.Itoa (i) + ".png"], false)
		time.Sleep (20 * time.Millisecond)
	}

	e.dead = true
}

func SpawnExplosion(pos sf.Vector2f) {
	exp := NewExplosion (pos)

	explosions = append (explosions, exp)
}