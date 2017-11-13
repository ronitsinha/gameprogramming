package main

import "fmt"

func main() {
	// Maps are like arrays that use types as indeces

	// map[*type of index*]*type of value*
	names := make (map[string]string)
	// Adds value "sinha" at "ronit"
	names ["david"] = "burton"

	//fmt.Println (names["david"])

	// Checks if map value exists and stores bool in ok, then if exists println
	if name, ok := names ["david"]; ok {
		fmt.Println (name)
	}

	for k, v := range names {
		fmt.Println (k, ":", v)
	}
}