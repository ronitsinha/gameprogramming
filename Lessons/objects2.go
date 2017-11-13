package main 

import "fmt"

type Circle struct {
	x float64
	y float64
	radius float64
}

type Rectangle struct {
	x float64
	y float64
	length float64
	width float64
}

// With input being in front, this function is only associated with circle objects
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Circumference () float64 {
	return 2 * c.radius * 3.14
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return r.length * 2 + r.width * 2
}

func main() {
	// All Circle objects will have the values that the object has (x, y, and radius)
	circle := &Circle {5, 0, 8}	

	// Calls function through an objects (functions like these are called "methods", in Go, functions attached to objects)
	a := circle.Area ()
	c := circle.Circumference ()

	rect := &Rectangle {5, 0, 1, 1}	

	// Calls function through an objects (functions like these are called "methods", in Go, functions attached to objects)
	a1 := rect.Area ()
	c1 := rect.Perimeter ()


	fmt.Println (a, c, a1, c1)

	//a := CircleArea (circle)

	// Does not specify a y value
	//circle0 := &Circle {x:5, radius:2}
}