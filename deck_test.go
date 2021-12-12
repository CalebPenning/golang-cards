package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 60 {
		t.Errorf("Expected deck length of 60 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in unshuffled deck to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Joker" {
		t.Errorf("Expected last card in unshuffled deck to be a Joker, but instead got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// damn that's a long test name!
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 60 {
		t.Errorf("Expected 60 cards in new deck, recieved %v", len(deck))
	}

	os.Remove("_decktesting")
}