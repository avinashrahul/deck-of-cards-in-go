package main

import "fmt"

// create new type that extends string properties

// below is read as type deck slice of string
type deck []string

// below d is a variable of type deck, usually it would be of 1 or 2 characters
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i)
		fmt.Println(card)
	}
}

// deck beside the newDeck function is that the return value should be of type deck

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Hearts"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

// a = [1, 2, 3, 4, 5]
// a[:2] => all elements from 0 to 2 except 2
// a[2:] => all elements from 2 to end
// a[0:3] => all elements from 0 to 3 except 3

func deal(d deck, deckSize int) (deck, deck) {
	return d[:deckSize], d[deckSize:]
}
