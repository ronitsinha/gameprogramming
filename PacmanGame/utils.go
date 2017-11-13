package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"time"
)

const (
	NODE_SIZE = 8 // Normally 8
	screenWidth  = 232
	screenHeight = 256
)

// Main
var res *Resources
var player *Player
var playerSpriteIndex int
var enemySpriteIndex int

// Maze
var gameMap *sf.Image
var wallSprites []*sf.Sprite
var grid [][]*Node
var pacdots []*sf.Sprite

// Ghosts
var ghosts []*Enemy

func Contains (array []*Node, element *Node) bool {
	for _, n := range array {
		if n == element {
			return true
		}
	}

	return false
}

func Reverse (array []*Node) {
	for i := len(array)/2-1; i >= 0; i-- {
		opp := len(array)-1-i
		array[i], array[opp] = array[opp], array[i]
	}
}

func NextGhost () {
	for _, g := range ghosts {
		if !g.isActive {
			g.isActive = true
			time.Sleep (10 * time.Second)
		}
	}
}