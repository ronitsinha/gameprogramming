package main

import (
	"os"
	"fmt"
)

func main() {
	// Prints out any inputted arguments
	//fmt.Println (os.Args)

	file, err := os.Open (".")

	// If there is an error, gtfo
	if err != nil {
		fmt.Println (err)
		return
	}

	
	// When creating a new variable using :=, only one needs to be new (files, not err)
	files, err := file.Readdir (0)
	if err != nil {
		fmt.Println (err)
		return
	}

	for i := 0; i < len(files); i ++  {
		fName := files[i].Name ()
		fmt.Println (fName)
	}

}