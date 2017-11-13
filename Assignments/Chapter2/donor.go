package main 

import "fmt"

func main() {
	var donation float64
	ranking := ""

	fmt.Print ("What is your donation? ")
	fmt.Scanf ("%f", &donation)

	if (donation >= 1000) {
		ranking = "benefactor"
	} else if (donation >= 500) {
		ranking = "patron"
	} else if (donation >= 100) {
		ranking = "supporter"
	} else if (donation >= 15) {
		ranking = "friend"
	} else {
		ranking = "cheapskate"
	}

	fmt.Println ("You are a", ranking + "!")
}