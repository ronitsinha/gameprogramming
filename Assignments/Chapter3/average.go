package main 

import "fmt"

func average(x []float64) float64 {
	var sum float64
	for i := 0; i < len(x); i++ {
		sum += float64(x[i])
	}
	sum /= float64(len (x))

	return sum
}

func main() {
	numbers := make ([]float64, 5)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = float64(i + 1)
	}

	fmt.Println (average(numbers))
}