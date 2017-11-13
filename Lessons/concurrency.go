package main

import (
	"fmt"
	"time"
)

// Channel filled with a certain type of channels (booleans in this case)
var done chan bool

// Concurrency: do multiple things at the same time (like run two functions at once)
// Threads: looks like doing things at the same time, but just switching between things quickly
// Goroutines in Go start a new thread

func foo(delay int) {
	for i := 0; i < 10; i ++ {
		fmt.Println (i)

		// Sleep pauses the program for a duration
		// Pauses for 1 second
		time.Sleep (time.Duration (delay) * time.Second)
	}

	// Stores a true in the channel once the function is completed
	done <- true
}

func main() {

	done = make (chan bool)

	// "go" will start its execution and then send it off to another thread, so things after it don't have to wait for it
	//go foo (1)
	//go foo (3)

	//var f func (int) = foo
	//go f (2)

	//defines a function
	/*var f func() = func() {
		fmt.Println ("Hello")
	}*/

	// defines and calls a function

	// these nameless functions are called "closuress"
	go func () {
		time.Sleep (5 * time.Second)
		fmt.Println ("Done waiting 5 seconds")
	}()

	// Once something is in the channel, it will be stored in this variable and removed from the channel
	// the variable will wait until the channel has a variable to set itself
	//x := <-done
	
	go foo (1)

	// removes the new value in the channel
	// Will see the value added by foo, remove it, and then the program will end
	<-done
	// Putting a go before this foo will end the program, because it will be passed off to another thread and continue
	//foo (3)
}

/*// This global variable can be used by any function (The level at which a variable can be accessed is its "scope")
var bar int

func foo() {
	bar ++
}

func main() {
	bar = 10

	foo ()
	foo ()

	fmt.Println (bar)
}*/