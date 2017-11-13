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
}


func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	grid := Grid {10, 10, Coord {0, 0}, Coord {0, 0}}

	playing := true

	correctPoints := make ([]Coord, 1)
	wrongPoints := make ([]Coord, 1)

	for (playing) {

		guessPoint := Coord {-1, -1}
		
		correctPoints [0] = Coord {-1, -1}
		wrongPoints [0] = Coord {-1, -1}

		grid.DisplayGrid (correctPoints, wrongPoints)

		fmt.Println (grid.actualPoint)
		fmt.Println (grid.badPoint)

		fmt.Println ("Guess an x coord")
		fmt.Scanf ("%d", &guessPoint.x)
		fmt.Println ("Guess a y coord")
		fmt.Scanf ("%d", &guessPoint.y)

		

		if (grid.CheckGuess (guessPoint) == 1) {
			correctPoints = append (correctPoints, guessPoint)
		}

		if (grid.CheckGuess (guessPoint) == 2) {
			wrongPoints = append (wrongPoints, guessPoint)
		}
	}

	fmt.Println ("END")
}

func (g *Grid) DisplayGrid(right, wrong []Coord) {

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


			for l := 0; l < len(right); l ++ {
				if j == right[l].x && i == right[l].y {
					gridRow[j] = "O"
				}
			}

			for k := 0; k < len(wrong); k ++ {
				if j == wrong[k].x && i == wrong[k].y {
					gridRow[j] = "X"
				}
			} 

			if i == randy && j == randx {
				g.actualPoint.x = j
				g.actualPoint.y = i
			} else if i == randy2 && j == randx2 {
				g.badPoint.x = j
				g.badPoint.y = i
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