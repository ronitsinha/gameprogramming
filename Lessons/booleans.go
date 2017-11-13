package main

import (
	"fmt"
)

func main() {
	x := false
	y := true
	z := x || y //x or y, if one is true, then this is true
	a := x && y //x and y, if both are true, then this is true
	b := !x // ! returns the opposite of the value (toggles it)

	fmt.Println (z, a, b)	
}