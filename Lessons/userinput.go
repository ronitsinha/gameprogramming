package main

import (
	"fmt"
)

func main() {
	fmt.Print ("Number, please: ") //Print does not create a new line after message, as opposed to Println

	var inputNum int

	const x = 10 // a constant is a value you can't change

	fmt.Scanf ("%d", &inputNum, x)
	// int input = %d
	// float64 input = %f

	fmt.Println (inputNum, "is a great number!")
}