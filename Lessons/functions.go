package main

import (
	"fmt"
)

// Declaring a function
// n int is the input, the final int determines the type of output
func Square(n int) int {
	return n*n
}

// if var type is on the last variable, like b, then the previous variables are assumed ot be of the same type
func Multiply(a, b int) (int, string) {
	return a*b, string(b*b)
}

func MultipleSquare(x, y int) (int, int) {
	return x*x, y*y
}

// the ...int indicates that "value" can take any number of integers
func Print(values ...int) {
	for i:= 0; i < len (values); i ++ {
		fmt.Print (values[i], " ")
	}
	
	// Just creates the new line
	fmt.Println ()
}

func Factorial(n int) int {
	// Factorial 0 is always 1
	if n <= 0 {
		return 1
	}
	// "Recursion" - When you call a function inside itself
	return n*Factorial(n-1)
}

func Compute(n int, computer func (int) int) int {
	return computer(n)
}

// The main function
func main() {
	/*var x int
	// "a" becomes 5^2, or 25, and "b" becomes 4^2, or 16
	a, b := MultipleSquare(5, 4)

	// 1 return for each Square func, so 2 needed for 2 vars
	//l, m := Square (5), Square(6)

	// This would not work, two vars but only 1 return
	// p, q := Square (10)
	// This would not work, since there are 2 returns and only one variable
	// r := MultipleSquare (2, 10)

	fmt.Print ("Enter an integer to be squared: ")
	fmt.Scanf ("%d", &x)

	fmt.Println (Square (x))
	fmt.Println (Multiply (6,8))
	fmt.Println (a, b)*/


	// Can take multiple values, sometimes called a "variatic" function
	//Print (1, 2, 5, 4, 10)
	//fmt.Println(Factorial (5))

	x := Factorial // x is the factorial function

	fmt.Println (x(5)) // same as factorial 5
	fmt.Println (Compute(5, x)) // uses func Factorial (x) to find the factorial of 5
}

