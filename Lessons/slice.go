package main

import (
	"fmt"
)

func main() {
	// Slice array, can be any length, no predetermined length
	//array[] int

	// Creates a "slice", which is an array with varying length
	// This is an int slice with 10 units of space
	/*array := make ([]int, 10)
	x := make ([]int, 9)

	// Appends a "15" to the end of the array, and makes the length that of the array you put int
	array = append (x, 15)

	// [2:9] returns elements 2 through 9, not including 9
	fmt.Println (array[2:9])
	
	// [:9] returns elements 0 through 9, not including 9
	fmt.Println (array[:9])

	// [2:] returns elements 2 until the end
	fmt.Println (array[2:])

	//*/

	// A string is a slice of runes
	//str := "Hello world" // Same type
	//strArray := make ([]rune, 11) // Same type

	// The []byte is a cast for the str variable 
	// len will return amount of bytes in string. A rune character can 
	//fmt.Println ([]byte (str), len (str))

	/*// Loops through an array, where 'i' is the index number (the place in the array) and n is the index value (the value stored in at the array's index)
	for _, n := range array {
		// do some stuff
	
	underscore '_' variables can be declared and not used
	}*/

}