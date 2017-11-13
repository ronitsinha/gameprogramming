package main 

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]

	SearchDir (".", string(fileName))
}

func SearchDir(dir, searchName string) {
	file, err := os.Open (dir)

	if err != nil {
		fmt.Println (err)
	}

	files, err := file.Readdir (0)

	if err != nil {
		fmt.Println ("BOO!", err)
		return
	}

	for i := 0; i < len (files); i ++ {
		if files[i].IsDir() {
			SearchDir (dir + "/" + files[i].Name(), searchName)
		} else if (files[i].Name () == searchName) {
			fmt.Println (dir + "/" + searchName)
		}
	}
}
