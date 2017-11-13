package main 

// This one will use "globs", also called regular expressions, for example *.go will find all .go files

import (
	"fmt"
	"os"
	"Lessons/glob"
)

func main() {
	fileName := os.Args[1]

	glob, _ := glob.Compile (os.Args[1])

	SearchDir (".", string(fileName))
}

func SearchDir(dir, searchName *glob.Glob) {
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
		} else if searchName.Match ([]byte (file.Name())) {
			fmt.Println (dir + "/" + searchName)
		}
	}
}
