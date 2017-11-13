package main 

import "fmt"

func main() {
	var array [100] int

	for i := 1; i < len(array); i++ {
		if (i % 3 == 0 && i % 5 == 0) {
			fmt.Println ("FizzBuzz")
		} else if (i % 3 == 0) {
			fmt.Println ("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println ("Buzz")
		} else {
			fmt.Println (i)
		}
	}
}