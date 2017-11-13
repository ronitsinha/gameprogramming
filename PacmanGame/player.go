package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"strconv"
	//"fmt"
)

type MoveKey struct {
	sf.KeyCode
	sf.Vector2f
	rotation float32
}

type Player struct {
	// embedding a sprite
	*sf.Sprite

	animateSpeed int
	// An array with a fixed length of 5
	keys [4]MoveKey
}

func NewPlayer(pos sf.Vector2f) *Player {
	player := new (Player)
	player.Sprite = sf.NewSprite (res.images["pacman_0.png"])
	player.SetOrigin (sf.Vector2f {player.GetGlobalBounds ().Width/2, player.GetGlobalBounds ().Height/2})
	player.animateSpeed = 0;

	player.SetPosition (pos)

	up := NewMoveKey (sf.KeyUp, sf.Vector2f{0, -1}, -90)
	down := NewMoveKey (sf.KeyDown, sf.Vector2f{0, 1}, 90)
	left := NewMoveKey (sf.KeyLeft, sf.Vector2f{-1, 0}, 180)
	right := NewMoveKey (sf.KeyRight, sf.Vector2f{1, 0}, 0)

	player.keys = [4]MoveKey {up, down, left, right}

	return player
}

func (p *Player) Update () {
	p.AnimatePlayer ((playerSpriteIndex/4)%3)

	if playerSpriteIndex >= 240 {
		playerSpriteIndex = 0
	} else {
		playerSpriteIndex ++
	}

	p.MovePlayer ()		
}


func NewMoveKey(key sf.KeyCode, move sf.Vector2f, rot float32) MoveKey {
	mK := MoveKey {key, move, rot}

	return mK
}

func (p *Player) AnimatePlayer(index int) {
	if p.animateSpeed != 0 {
		p.SetTexture (res.images["pacman_" + strconv.Itoa (index) + ".png"], false)
	} else {
		p.SetTexture (res.images["pacman_1.png"], false)
	}
}

func (p *Player) MovePlayer() {
	for _, k := range p.keys {
		if sf.KeyboardIsKeyPressed (k.KeyCode) {
			canMove := true
			canMinMove := false
			moveVector := k.Vector2f

			checkHeight := p.GetGlobalBounds().Height
			checkWidth := p.GetGlobalBounds().Width

			var checkRect sf.Rectf
			var checkMinRect sf.Rectf

			if moveVector.X < 0 {
				if moveVector.Y < 0 {
					checkRect = sf.Rectf {p.GetGlobalBounds().Left + moveVector.X, p.GetGlobalBounds().Top + moveVector.Y, checkWidth, checkHeight}
					checkMinRect = sf.Rectf {p.GetGlobalBounds().Left + 1, p.GetGlobalBounds().Top - 1, checkWidth, checkHeight}
				} else {
					checkRect = sf.Rectf {p.GetGlobalBounds().Left + moveVector.X, p.GetGlobalBounds().Top, checkWidth, checkHeight + moveVector.Y}
					checkMinRect = sf.Rectf {p.GetGlobalBounds().Left - 1, p.GetGlobalBounds().Top, checkWidth, checkHeight + 1}
				}
			} else {
				if moveVector.Y < 0 {
					checkRect = sf.Rectf {p.GetGlobalBounds().Left, p.GetGlobalBounds().Top + moveVector.Y, checkWidth + moveVector.X, checkHeight }
					checkMinRect = sf.Rectf {p.GetGlobalBounds().Left, p.GetGlobalBounds().Top - 1, checkWidth + 1, checkHeight -1}
				} else {
					checkRect = sf.Rectf {p.GetGlobalBounds().Left, p.GetGlobalBounds().Top, checkWidth + moveVector.X, checkHeight + moveVector.Y}
					checkMinRect = sf.Rectf {p.GetGlobalBounds().Left, p.GetGlobalBounds().Top, checkWidth + 1, checkHeight + 1}
				}
			}

			canMove = !CollisionWithLevel (checkRect)
			canMinMove = !CollisionWithLevel (checkMinRect)
			
			if canMove {
				p.SetRotation (k.rotation)
				p.Move (moveVector)
				p.animateSpeed = 1
			} else if canMinMove {
				p.SetRotation (k.rotation)
				p.Move (k.Vector2f)
				p.animateSpeed = 1
			}
		}
	}
	if !sf.KeyboardIsKeyPressed (sf.KeyUp) && !sf.KeyboardIsKeyPressed (sf.KeyDown) && !sf.KeyboardIsKeyPressed (sf.KeyLeft) && !sf.KeyboardIsKeyPressed (sf.KeyRight) {
		p.animateSpeed = 0
	}
}

func CollisionWithLevel (collider sf.Rectf) bool {

	for _, sp := range wallSprites {
		returnBool, _ := collider.Intersects (sp.GetGlobalBounds())
		if returnBool {
			return true
		}
	}

	return false
}