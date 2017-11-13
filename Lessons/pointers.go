package main 

import "fmt"

// The *int says the n is a memory address (the place where the memory is stored)
// n is a pointer
func f(n *int) {
	// x becomes 5 here
	// *n will find the value at the memory location and set it to 5
	// instead of setting variable, setting a value to a memory location, so the all copies of the variable that uses this memory location (x) will be changed
	*n = 6
	// Prints memory location
	fmt.Println (n)
}

func swap(a , b *int) {
	// *(variable) finds the value at a memory location
	/*tempA := *a
	*a = *b
	*b = tempA*/
	// More efficient/intuitive way to swap
	*a, *b = *b, *a
}

func main() {
	// Memory value of x will be copied and used when ever it is used again
	x := 10
	//& operator gets the memory location of the variable
	f (&x)
	// x is changed in f function (n = 5), but this is not changed
	// This is because the f function is changing a copy of x, using x's value, not the actual variable
	// Now that the value at x's memory location has been set to 5, this copy of x will also be 5 
	fmt.Println (x)

	a, b := 5, 10
	fmt.Println (a, b)
	swap (&a, &b)
	fmt.Println (a, b)
}