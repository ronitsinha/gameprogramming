package main 

import "fmt"

func main() {
	var age int
	var citizenYears int

	senate := false
	houseRep := false

	fmt.Print ("How old is the candidate? ")
	
	fmt.Scanf ("%d", &age)
	
	fmt.Print ("For how many years has the candidate been a U.S. citizen? ")
	
	fmt.Scanf ("%d", &citizenYears)

	if (age >= 25) {
		// Senate
		if (age >= 30 && citizenYears >= 9) {
			senate = true
		}
		// House of Reps
		if (citizenYears >= 7) {
			houseRep = true
		}
	}

	if (senate == true) {
		if (houseRep == true) {
			fmt.Println ("This candidate could be a Senator or a House Representative")
		} else {
			fmt.Println ("This candidate could be a Senator, but NOT a House Representative")
		}
	} else {
		if (houseRep == true) {
			// This case is not really necessary, as anyone who can be a Senator can be a House Rep
			fmt.Println ("This candidate could be a House Representative, but NOT a Senator")
		} else {
			fmt.Println ("This candidate could NOT be a Senator NOR a House Representative")
		}
	}
}