package main

func main() {
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards.saveToFile("test")
	cards.shuffle()
	cards.print()
}
