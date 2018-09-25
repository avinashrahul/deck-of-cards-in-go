package main

// := is initializing and assigning
func main() {
	cards := newDeck()

	first, last := deal(cards, 5)
	first.print()
	last.print()
}
