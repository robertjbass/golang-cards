package main

const FILEPATH = "my_cards.cards"

func main() {
	cards := newDeck()
	cards.saveToFile(FILEPATH)

	hand, cards := deal(cards, 5)
	cards.print()
	hand.print()

	cards = newDeckFromFile(FILEPATH)
	cards.shuffle()
	cards.print()
}
