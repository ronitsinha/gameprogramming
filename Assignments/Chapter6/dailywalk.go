package main

import (
	"fmt"
	"time"
	"math/rand"
)

var allready chan bool
var alarmDone chan bool
var alarmSet chan bool
var peopleReady int

func Ready(name string) {
	fmt.Println (name, "has started getting ready.")

	randInt := rand.Intn (20 - 10) + 10
	
	time.Sleep (time.Duration (randInt) * time.Second)
	
	fmt.Println (name, "spent", randInt, "seconds getting ready.")	
	peopleReady ++
}

func Alarm () {
	fmt.Println ("Arming alarm")
	alarmSet <- true
	time.Sleep (time.Duration (20) * time.Second)
	fmt.Println ("Alarm is armed")
	alarmDone <- true
}

func PuttingOnShoes(name string) {
	fmt.Println (name, "started putting on shoes.")

	randInt := rand.Intn (15 - 5) + 5
	
	time.Sleep (time.Duration (randInt) * time.Second)
	
	fmt.Println (name, "spent", randInt, "seconds putting on shoes.")
	peopleReady ++
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	allready = make (chan bool)
	alarmDone = make (chan bool)
	alarmSet = make (chan bool)

	bob := "Bob" 
	alice := "Alice"

	fmt.Println ("Let's go for a walk!")

	time.Sleep (time.Second)

	go func() {
		go Ready (alice)
		Ready (bob)
		if (peopleReady >= 2) {
			allready <- true
		}
	} ()

	<- allready
	go Alarm ()
	<- alarmSet

	go func() {
		peopleReady = 0
		
		go PuttingOnShoes (alice)
		PuttingOnShoes (bob)
		if (peopleReady >= 2) {
			allready <- true
		}
	} ()

	<- allready

	fmt.Println ("Exiting and locking door")

	<- alarmDone
}