// The largest number in decimal that a 10 bit integer can store is 1023

package main 

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print ("Enter a number of bits: ")
	var input float64
	var result float64

	fmt.Scanf ("%f", &input)

	result = math.Pow(2.0, input) - 1.0

	fmt.Println ("Decimal form:", result)

}

/*The maximum value of 32 bit integer is 2,147,483,647. This has to do with the fact that it is a signed value, and signed values
also have representations for negative numbers, so basically a max 32 signed int would be (2^31)-1, with 31 instead of 32 because 1
bit is used for positive or negative interpretations. This expression simplifies to 2,147,483,647, instead of 4,294,967,295, which is 
what one would expect the max value would be*/