package main 

import "fmt"

func fibR(n int) int {
	if (n == 0) {
		return 0
	} else if (n == 1) {
		return 1
	}

	return fibR (n-1) + fibR (n-2)
}

func fibIter(n int) int {
	num1 := 0
	num2 := 1
	var fibonacci int
	for i := 0; i < n; i++ {
		fibonacci = num1 + num2
		num1 = num2
		num2 = fibonacci
	}
	return num1
}

func main() {
	//fmt.Println (fibR(40))
	fmt.Println (fibIter(40))
}