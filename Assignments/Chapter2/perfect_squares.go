package main 

import (
	"fmt"
	"math"
)

func main() {
	perfectSquareCount := 0
	var num int

	fmt.Print ("Enter an integer: ")
	fmt.Scanf ("%d", &num)


	for i := 1; i < num; i ++ {
		sqrtI := math.Pow (float64(i), 0.5)
		if sqrtI == float64 (int64(sqrtI)) {
			perfectSquareCount ++
			fmt.Println (i)
		}
	}

	fmt.Println ("There is/are", perfectSquareCount, "perfect square/s between 1 and", num)

}