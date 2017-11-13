package main 

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	radius float64
}

func CircleArea(c *Circle) float64 {
	return (math.Pi * c.radius * c.radius)
}

func MakeUnitCircle() *Circle {
	circle := &Circle {0, 0, 1}
	return circle
}

func main() {
	// Prints out value for Pi
	fmt.Println (CircleArea(MakeUnitCircle()))
}