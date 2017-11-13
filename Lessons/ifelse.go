package main 

import "fmt"

func main() {
	var input int

	fmt.Scanf ("%d", &input)

	// Prints "Hello world" if value is true
	if input == 10 {
		fmt.Println ("You entered 10")
	// Runs if the if statement doesn't
	} else if input == 11{
		fmt.Println ("You entered 11")
	// Runs if the previous two statements don't
	} else {
		fmt.Println ("You did not enter 10 or 11")
	}

	fmt.Println ("Hello World")
}