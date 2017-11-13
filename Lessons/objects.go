// Chapter 9 in textbook

 package main

import (
	"fmt"
	"math"
)

// Can make your own types (ie int) called objects or structs
// This is a new data type (type of data)
type bite byte

type Point struct {
	x float64
	y float64
}

// Can find distance between 2 points
func Distance (p1, p2 Point) float64 {
	return math.Sqrt (math.Pow (p1.x-p2.x, 2) + math.Pow (p1.y-p2.y, 2))
}

// From now on, ALWAYS use a pointer for objects
func Zero(p1 Point) {
	// Data types always use pointers
	// Really is *p1.x
	p1.x = 0
	// Really is *p1.y
	p1.y = 0
}

func main() {
	// Same as 'var x byte'
	//var x bite

	// Reads variables down the list (first will be x, second will be y)
	p := Point {10, 5}
	p1 := Point {10, 5}

	// You can actually change the order in which variables are assigned
	//p := Point {y:10, x:5}

	// Prints x value of Point p (10)
	fmt.Println (p.x)
	// Prints y value of Point p (5)
	fmt.Println (p.y)


	fmt.Println (Distance(p, p1))
}