package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card is King of Diamonds, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const TEST_FILENAME = "_decktesting.cards"
	os.Remove(TEST_FILENAME)
	deck := newDeck()

	err := deck.saveToFile(TEST_FILENAME)
	if err != nil {
		t.Errorf("Expected to write cards to filesystem but received a error - %v", err)
	}

	loadedDeck := newDeckFromFile(TEST_FILENAME)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 51 cards in loadedDeck, but got %v", len(loadedDeck))
	}

	os.Remove(TEST_FILENAME)
}
