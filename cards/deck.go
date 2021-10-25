package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck', which is a slice of strings
type deck []string

// (d deck) is a reciever
// Any variable inside our application of type 'deck' gets access to this function (the 'print' method)

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Creates a new Deck 'deck' and returns it as a 'deck' type
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

// Takes a deck and joins each element into a string-based CSV.
// Returns the joined strings as one massive string, comman-separated
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0660)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()

		// Option 2 - log the error and entirely quit the program
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}
