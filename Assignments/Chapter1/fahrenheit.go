package main

import (
	"fmt"
)

func main(){
	var f float32
	var c float32
	
	fmt.Println("Enter a temperature in Fahrenheit: ")
	
	fmt.Scanf("%f",&f)
	
	c = (f-32)*5/9
	
	fmt.Println(f, "degrees Fahrenheit is", c , "degrees Celcius")
}