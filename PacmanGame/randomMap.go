package main 

import (
	sf "github.com/zyedidia/sfml/v2.3/sfml"
	"math/rand"
	"math"
	"time"
)

const (
	CELL_WIDTH = 16
	CELL_HEIGHT = 16

	VOID = 1000
	WALL = 100
	FLOOR = 1
)

type Grid struct {
	allCells [][]int
}


func NewGrid (w, h int) *Grid {
	rgrid := new (Grid)
	rgrid.allCells = make ([][]int, w)

	for x := 0; x < w; x ++ {
		cellListY := make ([]int, h)
		for y := 0; y < h; y ++ {
			cellListY [y] = VOID
		}	
		rgrid.allCells [x] = cellListY 
	}

	return rgrid
}

func (g *Grid) SetRegion(x1, y1, x2, y2, val int) {
	for y := y1; y < y2; y ++ {
		for x := x1; x < x2; x ++ {
			g.allCells[x][y] = val
		}
	}
}



func Generate () []*sf.RectangleShape {
	//50
	width := int (screenWidth/CELL_WIDTH)
	//37
	height := int (screenHeight/CELL_HEIGHT)


	returnArray := make ([]*sf.RectangleShape, 0)

	//Grid
	grd := NewGrid (width, height)

	//Fill Grid
	grd.SetRegion (0, 0, width-1, height-1, VOID)

	//Random Seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Controller
	cx := int (width/2)
	cy := int (height/2)

	//Controller Direction
	cdir := rand.Intn(3)

	//Odds
	odds := 1

	//Create Level
	for i := 0; i < 1000; i ++ {
		//Place floor
		grd.allCells[cx][cy] = FLOOR

		rand.Seed(time.Now().UTC().UnixNano())
		//Randomize Directions
		if (rand.Intn(odds) == odds) {

			cdir = rand.Intn(3)
		}

		//Move
		xdir := math.Cos (float64 (cdir*90)) * float64 (rand.Intn(3) + 1)
		ydir := math.Sin (float64 (cdir*90)) * float64 (rand.Intn(3) + 1)

		cx += int (xdir)
		cy += int (ydir)

		//Don't move outside
		if cx > width-2 {
			cx = width-2
		} else if cx < 1 {
			cx = 1
		}

		if cy > height-2 {
			cy = height-2
		} else if cy < 1 {
			cy = 1
		}
	}

	for yy := 1; yy < height-1; yy ++ {
		for xx := 1; xx < width-1; xx ++ {
			if grd.allCells [xx][yy] == FLOOR {
				if grd.allCells [xx+1][yy] != FLOOR {grd.allCells [xx+1][yy] = WALL}
				if grd.allCells [xx-1][yy] != FLOOR {grd.allCells [xx-1][yy] = WALL}
				if grd.allCells [xx][yy + 1] != FLOOR {grd.allCells [xx][yy + 1] = WALL}
				if grd.allCells [xx][yy - 1] != FLOOR {grd.allCells [xx][yy - 1] = WALL}
			}
		}
	}
	//Draw the level
	for yy := 0; yy < height; yy ++ {
		for xx := 0; xx < width; xx ++ {
			if grd.allCells [xx][yy] == FLOOR {
				rect := sf.NewRectangleShape (sf.Vector2f {float32(CELL_WIDTH), float32(CELL_HEIGHT)})
				rect.SetPosition (sf.Vector2f {float32(xx*CELL_WIDTH), float32(yy*CELL_HEIGHT)})
				rect.SetFillColor (sf.ColorBlue)
				returnArray = append (returnArray, rect)
			} else if grd.allCells [xx][yy] == WALL {
				rect := sf.NewRectangleShape (sf.Vector2f {float32(CELL_WIDTH), float32(CELL_HEIGHT)})
				rect.SetPosition (sf.Vector2f {float32(xx*CELL_WIDTH), float32(yy*CELL_HEIGHT)})
				rect.SetFillColor (sf.ColorBlack)
				returnArray = append (returnArray, rect)
			}
		}
	}

	return returnArray	
}