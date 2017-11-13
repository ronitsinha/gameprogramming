package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	phrase := os.Args[1]

	FindFileWithPhrase (".", phrase)

}

func FindFileWithPhrase(dir, keyPhrase string) {
	file, err := os.Open (dir)

	if err != nil {
		fmt.Println (err)
		return
	}

	files, err := file.Readdir (0)

	if err != nil {
		fmt.Println ("BOO!", err)
		return
	}

	for i := 0; i < len (files); i ++ {

		if files[i].IsDir () {
			FindFileWithPhrase (dir + "/" + files[i].Name(), keyPhrase)
		} else {
			fileReadInit, err := ioutil.ReadFile (dir + "/" + files[i].Name())

			if err != nil {
				fmt.Println (err)
				return
			}

			fileRead := strings.Split (string (fileReadInit), "\n")

			for j := 0; j < len (fileRead); j ++ {
				if strings.Contains(fileRead[j], keyPhrase) {
					fmt.Println ("In " + dir + "/" + files[i].Name() + ":")
					fmt.Println (keyPhrase, "on line", j + 1, ">>", fileRead[j])
				}
			}
		}
	}
}