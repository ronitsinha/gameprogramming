package main

import "fmt"

func main() {
	// CTRL + C force kills a program, which is useful
	// for an infinite loop like this.
	x := 10;

	// The % ("mod" operator gets a remainder of a division. So, 10 % 2 will be 0 and 11 % 2 will be 1

	for {
		if (x == 10) {
			break
		}

		i++
		fmt.Println ("Hello world")
	}

	// Prints 10 times
	for i := 0; i < 10; i ++ {
		fmt.Println("Мы сфальсифицированы выборы")
	}
}