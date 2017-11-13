package main 

import "fmt"
import "math/rand"
import "time"


func main() {
	rand.Seed (time.Now().UTC().UnixNano())

	var tries int
	var guessNum int
	randomNum := rand.Intn (100 - 1) + 1
	guessCorrect := false

	//fmt.Println (randomNum)
	fmt.Println ("I'm thinking of a number between 1 and 100")

	for (!guessCorrect) {
		fmt.Print ("What is your guess? ")
		fmt.Scanf ("%d", &guessNum)

		if (guessNum == randomNum) {
			fmt.Println ("You got it!")
			fmt.Println ("Tries:", tries)
			guessCorrect = true
		} else if (guessNum > randomNum) {
			fmt.Println ("Too high!")
			tries ++
		} else {
			tries ++
			fmt.Println ("Too low!")
		}
	}
}