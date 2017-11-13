package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math/rand"
)

type Node struct {
	*sf.RectangleShape

	walkable bool
	gridX int
	gridY int

	gCost int
	hCost int
	parent *Node
}

func (n *Node) fCost () int {
	return n.gCost + n.hCost
}

func NewNode (pos sf.Vector2f, gX, gY int) *Node {
	nde := new (Node)

	nde.RectangleShape = sf.NewRectangleShape (sf.Vector2f {float32 (NODE_SIZE), float32 (NODE_SIZE)})
	nde.SetPosition (pos)
	nde.gridX = gX
	nde.gridY = gY

	nde.walkable = true

	return nde
}

func (nde *Node) FindAdjacentWalkables () []*Node {
	nodeSlice := make ([]*Node, 0)

	for x := -1; x <= 1; x ++ {
		for y := -1; y <= 1; y ++ {
			if x == y || x == -y {
				continue
			}

			checkX := nde.gridX + x
			checkY := nde.gridY + y

			if (checkX >= 0 && checkX < len(grid)) && (checkY >= 0 && checkY < len(grid[0])) {
				if grid[checkX][checkY].walkable {
					nodeSlice = append (nodeSlice, grid[checkX][checkY])
				}
			}
		}
	}

	return nodeSlice	
}

func LoadMap () {
	gameMap = sf.NewImage ("./assets/images/maze.png")

	for y := 0; y <	int(gameMap.GetSize ().Y); y ++ {
		for x := 0; x <	int(gameMap.GetSize ().X); x ++ {
			if gameMap.GetPixel (uint(x), uint(y)).IsEqual (sf.ColorBlue) {
				appendSprite := sf.NewSprite (sf.NewEmptyTexture (1, 1))
				appendSprite.SetPosition (sf.Vector2f {float32 (x), float32(y)})

				wallSprites = append (wallSprites, appendSprite)
			}
		}
	}

	NodeGrid ()
	GeneratePacDots ()
}

func NodeFromPosition (pos sf.Vector2f) *Node {
	for x, _ := range grid {
		for _, nde := range grid[x] {
			if nde.GetGlobalBounds ().Contains (pos) {
				return nde
			}
		}
	}

	return nil
}

func NodeGrid () {
	grid = make ([][]*Node, 0)

	for x := 0; x <	int(gameMap.GetSize ().X/uint (NODE_SIZE)); x ++ {
		gridY := make ([]*Node, 0)
		for y := 0; y <	int(gameMap.GetSize ().Y/uint (NODE_SIZE)); y ++ {
			addNode := NewNode (sf.Vector2f {float32 (x * NODE_SIZE), float32(y * NODE_SIZE)}, x, y)
			for _, spr := range wallSprites {
				inWall, _ := spr.GetGlobalBounds ().Intersects (addNode.GetGlobalBounds ()) 
				if inWall {
					addNode.walkable = false
				}
			}

			gridY = append (gridY, addNode)
		}
		grid = append (grid, gridY)
	}
}


func GeneratePacDots () {
	pacdot := res.images ["pacdot.png"]
	pacball := res.images ["pacball.png"]

	for x, _ := range grid {
		for _, n := range grid[x] {
			if n.walkable {
				if rand.Intn (10) == 4 {
					pball := sf.NewSprite (pacball)
					pball.SetOrigin (sf.Vector2f {pball.GetGlobalBounds ().Width/2, pball.GetGlobalBounds ().Height/2})
					pball.SetPosition (sf.Vector2f {n.GetPosition ().X + NODE_SIZE/2, n.GetPosition ().Y + NODE_SIZE/2})

					pacdots = append (pacdots, pball)
				} else {
					pdot := sf.NewSprite (pacdot)
					pdot.SetOrigin (sf.Vector2f {pdot.GetGlobalBounds ().Width/2, pdot.GetGlobalBounds ().Height/2})
					pdot.SetPosition (sf.Vector2f {n.GetPosition ().X + NODE_SIZE/2, n.GetPosition ().Y + NODE_SIZE/2})

					pacdots = append (pacdots, pdot)
				}
			}
		}
	}
}