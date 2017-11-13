package main 

import "fmt"

func main() {
	// this array called "array" will store a list of 4 integers
	var array [10000] int
	//array [0] = 5
	for i := 0; i < len (array); i ++ {
		array [i] = i;
	}

	// the & operator finds the memory location where a variable is stored

	fmt.Print (array)

}