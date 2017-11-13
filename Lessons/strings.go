package main

import (
	"fmt"
	//"strconv" // String conversions
)

func main() {
	x := 65;

	y := string (x) // converts int x to string y 
	// conversions are called casting


	fmt.Println (y, x);

	//var x string = "Hello World" // a string is a string of runes and bytes
	//fmt.Println (x)
	//fmt.Println (len(x)) // prints the length (amount of bytes) in a string
}