package main

import (
	"fmt"
)

func main() {
	// var card string = "Hello world"

	card := newCard() // both are same, assigning data to a variable
	// card = "Hello guru"   // re-assign variable declaration

	fmt.Println(card)

	cards := deck{newCard(), "new array of cards"}

	deck.print(cards)

}

func newCard() string {
	return "returned string"
}
