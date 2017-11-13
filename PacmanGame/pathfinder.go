package main

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math"
	//"fmt"
)

func FindPath(startPos, targetPos sf.Vector2f) []*Node {
	startNode := NodeFromPosition (startPos)
	targetNode := NodeFromPosition (targetPos)

	openSet := make ([]*Node, 0)
	closeSet := make ([]*Node, 0)
	openSet = append (openSet, startNode)

	for len (openSet) > 0 {
		currentNode := openSet[0]

		index := 0

		for i, n := range openSet {
			if n.fCost () < currentNode.fCost () || (n.fCost () == currentNode.fCost () && n.hCost < currentNode.hCost) {
				currentNode = n
				index = i
			}
		}

		copy (openSet[index:], openSet[index+1:])
		openSet[len(openSet) - 1] = nil
		openSet = openSet[:len(openSet)-1]

		closeSet = append (closeSet, currentNode)

		if currentNode == targetNode {
				return RetracePath (startNode, targetNode)
		}

		for _, adj := range currentNode.FindAdjacentWalkables () {
			if Contains (closeSet, adj) {
				continue
			}

			skip := false

			for _, g := range ghosts {
				if adj.GetGlobalBounds ().Contains (g.GetPosition ()) {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			newMovementCostToAdj := currentNode.gCost + GetDistance (currentNode, adj)
			if newMovementCostToAdj < adj.gCost || !Contains (openSet, adj) {
				adj.gCost = newMovementCostToAdj
				adj.hCost = GetDistance (adj, targetNode)
				adj.parent = currentNode

				if !Contains (openSet, adj) {
					openSet = append (openSet, adj)
				}
			}
		}
	}

	return nil
}

func RetracePath (start, end *Node) []*Node {
	currentNode := end
	path := make ([]*Node, 0)

	for currentNode != start {
		path = append (path, currentNode)
		currentNode = currentNode.parent
	}

	Reverse (path)

	/*isPath := false

	for _, n := range path {
		if Contains (start.FindAdjacentWalkables (), n) {
			isPath = true
		}
	}

	if isPath {
		fmt.Println ("The path starts next to me!")
	}*/

	return path
}


func GetDistance (nodeA, nodeB *Node) int {
	distX := int (math.Abs (float64 (nodeA.gridX - nodeB.gridX)))
	distY := int (math.Abs (float64(nodeA.gridY - nodeB.gridY)))

	if distX > distY {
		return 14*distY + 10*(distX - distY)
	}

	return 14*distX + 10*(distY - distX)
}