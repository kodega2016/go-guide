package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected total length of the cards is 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first element of the deck is Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last element of the deck is Four of Clubs but got %s", d[len(d)-1])
	}
}

func TestNewDeckAndSaveDeckToFile(t *testing.T) {
	os.Remove("__deck_testing")
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected total length of the cards is 16 but got %v", len(d))
	}
	d.saveDeckToFile("__deck_testing")
	deck_from_file, _ := newDeckFromFile("__deck_testing")
	if len(deck_from_file) != 16 {
		t.Errorf("Expected length of the cards is 16 but got %v", len(deck_from_file))
	}
	os.Remove("__deck_testing")
}

