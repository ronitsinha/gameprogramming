package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	wordArray := []string {"apple", "bike", "cat", "dinosaur", "eggplant"}

	chosenWord := wordArray [random (0, len(wordArray))]
	stillPlaying := true
	var guessLetter rune
	hiddenWord := make ([]rune, len(chosenWord))

	for i:= 0; i < len (hiddenWord); i++ {
		hiddenWord[i] = '-'
	}

	for (stillPlaying) {
		correct := false;
		revealedLetters := 0
		wrongLetters := 0

		fmt.Println (string(hiddenWord)) 
		fmt.Print ("What's your guess? ")
		fmt.Scanf ("%c\n", &guessLetter)
		
		for i:= 0; i < len (chosenWord); i++ {

			chosenRuneArray := []rune(chosenWord)

			if guessLetter == chosenRuneArray[i] {
				correct = true
				hiddenWord[i] = guessLetter
			} else {
				wrongLetters ++
			}
		
			if hiddenWord[i] != '-' {
				revealedLetters ++
			}
		}

		if (wrongLetters == len(hiddenWord)) {
			fmt.Println ("Incorrect")
		}

		if (revealedLetters == len(hiddenWord)) {
			fmt.Println (string(hiddenWord)) 
			stillPlaying = false
		}

		if correct {
			fmt.Println ("Correct!")
			correct = false
		}
	}

	fmt.Println ("Yes! The word was", chosenWord + "!")

}