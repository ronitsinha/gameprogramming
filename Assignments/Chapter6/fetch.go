package main 

import (
	"fmt"
	"time"
	"net/http"
	"os"
)

func fetch(url string, ch chan string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error: %v\n", err)
        return
    }
    resp.Body.Close()
    ch <- fmt.Sprintf("%.2fs %s\n", time.Since(start).Seconds(), url)
}

func main() {
	var ch chan string = make (chan string)

	if len (os.Args) >= 2 {
		for i := 1; i < len(os.Args); i++ {
			go fetch (os.Args[i], ch)
			x := <-ch
			fmt.Println (x)
		}
	}

	fmt.Println ("Done")
}