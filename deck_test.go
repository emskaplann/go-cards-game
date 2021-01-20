package main

import (
	"math/rand"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()

	// Check length
	if len(newDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d.", len(newDeck))
	}

	// Check first el
	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got %s.", newDeck[0])
	}

	// Check last el
	if newDeck[len(newDeck)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of 'Four of Clubs', but got %s.", newDeck[len(newDeck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	randomNumber := rand.Intn(len(deck) - 1)
	if deck[randomNumber] != loadedDeck[randomNumber] {
		t.Errorf("Created and loaded decks are not same.")
	}
	os.Remove("_decktesting")
}
