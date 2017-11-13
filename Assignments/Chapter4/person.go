package main 

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	age int
	emailAddress string
}

func Describe(p *Person) {
	fmt.Print (p.firstName + " " + p.lastName, " is ", p.age, " years old. For more information, you can contact ", p.firstName, " at ", p.emailAddress)	
}

func main() {
	john := &Person {"John", "Smith", 27, "jsmith@gmail.com"}

	Describe (john)
}