package main 

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]

	//fmt.Println ("Enter file name:")
	//fmt.Scanf ("%s", &fileName)

	contents, err := ioutil.ReadFile (fileName)

	fmt.Println ()

	if (err != nil) {
		fmt.Println (err)
		return
	}

	fmt.Println (string(contents))
}

