package main

import (
	"fmt"
)

type Color struct {
	r, g, b byte
}

type Vehicle struct {
	maxSpeed int
	brand string
	color Color
}

func (v *Vehicle) Repaint(newColor Color) {
	// Changes a vehicle's, even embedded ones like Car, color
	v.color = newColor
}

// Has all vehicle qualities, but also more
type Car struct {
	// Just the type to "inherit" from
	// * for types, & for variables
	*Vehicle

	seats int
	plateNum string

}

func main() {
	// Put in a new vehicle to satifsy inheritance, called struct embedding in Go, of Car from Vehicle
	car := &Car {&Vehicle{100, "McQueen", Color {255, 0, 0}}, 2, "K4CH0W"}

	// Directly access Vehicle values from the car
	fmt.Println (car.color)

	car.Repaint (Color {0, 0, 255})
	fmt.Println (car.color)
}
