package main 

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

type Coord struct {
	x int
	y int
}

type Grid struct {
	gridL int
	gridW int
	actualPoint Coord
	badPoint Coord
	openPoints []Coord
}


func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	grid := Grid {5, 5, Coord {0, 0}, Coord {0, 0}, nil}

	playing := true
	endgame := false
	openFilled := false

	correctPoints := make ([]Coord, 1)
	wrongPoints := make ([]Coord, 1)
	guessedPoints := make ([]Coord, 1)

	for (playing) {

		guessPoint := Coord {-1, -1}
		
		correctPoints [0] = Coord {-1, -1}
		wrongPoints [0] = Coord {-1, -1}
		guessedPoints[0] = Coord {-1, -1}

		grid.DisplayGrid (correctPoints, wrongPoints, guessedPoints, openFilled, endgame)
		openFilled = true


		if (endgame) {
			playing = false
			break
		}

		//fmt.Println (grid.openPoints)


		//fmt.Println (grid.actualPoint)
		//fmt.Println (grid.badPoint)

		fmt.Println ("Guess an x coord")
		fmt.Scanf ("%d", &guessPoint.x)
		fmt.Println ("Guess a y coord")
		fmt.Scanf ("%d", &guessPoint.y)


		endgame = grid.RemoveOpenPoint (guessPoint)
		

		if (grid.CheckGuess (guessPoint) == 1) {
			correctPoints = append (correctPoints, guessPoint)
		} else if (grid.CheckGuess (guessPoint) == 2) {
			wrongPoints = append (wrongPoints, guessPoint)
			endgame = true
		} else {
			guessedPoints = append (guessedPoints, guessPoint)
		}
	}

	fmt.Println ("END")
}

func (g *Grid) DisplayGrid(right, wrong []Coord, guessed []Coord, openBool bool, gameEnd bool) {

	var spaceNum int

	randx := rand.Intn (g.gridW)
	randy := rand.Intn (g.gridL)

	randx2 := rand.Intn (g.gridW)
	randy2 := rand.Intn (g.gridL)

	if g.gridL < 10 {
		spaceNum = 2
	} else if g.gridL < 100 {
		spaceNum = 1
	} else if g.gridL < 1000 {
		spaceNum = 0
	}

	gridRow := make ([]string, g.gridW)

	rowCount := make ([]int, len(gridRow))

 	for i := 0; i < len(gridRow); i ++ {
 		rowCount[i] = i

 	}

	for n := 0; n < spaceNum + 1; n ++ {
		fmt.Print (" ")
	}

 	fmt.Println (rowCount)

	for i := 0; i < g.gridL; i ++ {
	
		for j := 0; j < len (gridRow); j ++ {
			
			gridRow[j] = "-"

			if !openBool {
				g.openPoints = append (g.openPoints, Coord {j, i})
			}

			for l := 0; l < len(right); l ++ {
				if j == right[l].x && i == right[l].y {
					gridRow[j] = "O"
					//gameEnd = g.RemoveOpenPoint (Coord {j, i})
				}
			}

			for k := 0; k < len(wrong); k ++ {
				if j == wrong[k].x && i == wrong[k].y {
					gridRow[j] = "X"
					//gameEnd = g.RemoveOpenPoint (Coord {j, i})
				}
			} 

			for z := 0; z < len(guessed); z ++ {
				if j == guessed[z].x && i == guessed[z].y {
					gridRow[j] = " "
					//gameEnd = g.RemoveOpenPoint (Coord {j, i})
				}
			} 

			if (g.IsOpenPoint (Coord {j, i})) {
				if i == randy && j == randx {
					g.actualPoint.x = j
					g.actualPoint.y = i
					
					randx = rand.Intn (g.gridW)
					randy = rand.Intn (g.gridL)

				} else if i == randy2 && j == randx2 {
					g.badPoint.x = j
					g.badPoint.y = i
					
					randx2 = rand.Intn (g.gridW)
					randy2 = rand.Intn (g.gridL)
				}
			} else {
				gameEnd = true
			}
		}
		fmt.Print (strconv.Itoa(i))

		for n := 0; n < spaceNum; n ++ {
			fmt.Print (" ")
		}

		fmt.Println (gridRow)
	}
}

func (g *Grid) CheckGuess(guess Coord) int {

	if (guess.x == g.actualPoint.x && guess.y == g.actualPoint.y) {
		fmt.Println ("Correct!")
		return 1
	} else if (guess.x == g.badPoint.x && guess.y == g.badPoint.y) {
		fmt.Println ("Uh-oh!")
		return 2
	}

	fmt.Println ("Nothing there!")
	return 3
}

func (g *Grid) IsOpenPoint (point Coord) bool {
	for i := 0; i < len (g.openPoints); i ++ {
		if point.x == g.openPoints[i].x && point.y == g.openPoints[i].y {
			return true
		}
	}

	return false
}

func (g * Grid) RemoveOpenPoint (point Coord) bool {
	var indexToRemove int

	for i := 0; i < len (g.openPoints); i ++ {
		if point.x == g.openPoints[i].x && point.y == g.openPoints[i].y {
			indexToRemove = i
		}
	}

	if (len (g.openPoints) >= 2) {
		g.openPoints = append (g.openPoints[:indexToRemove], g.openPoints[indexToRemove+1:]...)
		return false
	}

	return true
}