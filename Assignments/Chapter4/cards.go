package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

type Card struct {
	suit int
	value int
	suitName string
	valueName string
}

func MakeDeck() []Card {
	// Make 4 slices to sort cards by suit
	clubs := make ([]Card, 13)
	spades := make ([]Card, 13)
	hearts := make ([]Card, 13)
	diamonds := make ([]Card, 13)

	// This array (2D) is a collection of the suit slices
	masterArray := [][]Card {clubs, spades, hearts, diamonds}

	// This will be the returned slice
	cards := make ([]Card, 0)

	// This for loop goes through the indexes in the suit slices using their indexes in the master array
	for i := 0; i < len(masterArray); i ++ {
		for j := 0; j < 13; j ++ {
			masterArray[i][j].suit = i
			masterArray[i][j].value = j+1

			masterArray[i][j].valueName = FindValue (masterArray[i][j].value)
			masterArray[i][j].suitName = FindSuit (masterArray[i][j].suit)

			cards = append (cards, masterArray[i][j])
		}
	}

	return cards
}

func FindSuit (suit int) string {
	if suit == 0 {
		return "CLUBS"	
	} else if suit == 1 {
		return "SPADES"
	} else if suit == 2 {
		return "HEARTS" 
	} else {
		return "DIAMONDS"
	}
}

func FindValue(value int) string {

	if value == 1 {
		return "Ace"
	} else if value == 11 {
		return "Jack"
	} else if value == 12 {
		return "Queen"
	} else if value == 13 {
		return "King"
	} else {
		return strconv.Itoa (value)
	}

	return strconv.Itoa (value)
}

func PickRandomCard(cardArray []Card) Card {
	indexToPick := rand.Intn (len (cardArray))
	pickedCard := cardArray[indexToPick]

	cardArray = append (cardArray[:indexToPick], cardArray[indexToPick+1:]...)

	return pickedCard
}

func Shuffle(cardArray []Card) []Card {
	shuffledArray := make ([]Card, len(cardArray))
	perm := rand.Perm (len (cardArray))

	for i := 0; i < len (perm); i ++ {
		shuffledArray[i] = cardArray[perm[i]]
	}

	return shuffledArray
}

func PrintCard (c Card) string {
	str := c.valueName + " of " + c.suitName + ", "

	return str
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	deck := MakeDeck()

	// Shuffles the deck
	deck = Shuffle(deck)

	//fmt.Println (deck)
	for i := 0; i < len (deck); i ++ {
		fmt.Print (PrintCard(deck[i]))
	}
	// Picks a random card from the deck and removes it
	fmt.Println ()
	fmt.Println ()
	fmt.Println (PrintCard(PickRandomCard (deck)))
	fmt.Println ()
	//fmt.Println(deck)
	for i := 0; i < len (deck); i ++ {
		fmt.Print (PrintCard(deck[i]))
	}
}