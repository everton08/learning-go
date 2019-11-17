package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	deckLength := 16
	d := newDeck()

	if len(d) != deckLength {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_pepo_cards_test"
	deckLength := 16

	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != deckLength {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	os.Remove(fileName)
}