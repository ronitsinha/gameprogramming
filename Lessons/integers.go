package main

import (
	"fmt"
)

func main() {
	var x uint64 = 10 // unsigned int, no bit for negative or positive distinguishment, 64 bits of storage
	var y int64 = 5 // signed int, 63 bits of digit storage, with 1 bit being used for a negative or positive distinguishment
 	// an 8-bit int can store an 8 digit int

 	var z float64 = 1.5e+10 // can store more because use certain number of bits for before and after decimal point, and for exponent 

 	// an unsigned 8-bit integer (uint8) is a byte
 	var w byte = 'A' // single quotes can have only one character


 	var p rune = 'ε' // unicode (more storage) byte
 	// Runes can store characters with higher values

	fmt.Println (x, y, z, w, p)
}