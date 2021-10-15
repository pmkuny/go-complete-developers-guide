package main

import "fmt"

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

func Deal(d deck, c int) deck {

}
