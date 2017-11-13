package main

import (
	"fmt"
	"time"
)

func Repeat(f func(), duration time.Duration) {
	go func () {	
		for {
			f ()	
			time.Sleep (duration)
		}
	}()
}

func SayHello() {
    fmt.Println("Hello")
}

func SayGoodbye() {
    fmt.Println("Goodbye")
}

func main() {
	Repeat(SayHello, 2 * time.Second)
    Repeat(SayGoodbye, 3 * time.Second)

    time.Sleep(10 * time.Second)
}