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
