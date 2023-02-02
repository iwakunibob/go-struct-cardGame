package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0] != "2♦" {
		t.Errorf("Expected first card 2♦ but got %v", d[0])
	}
	if d[len(d)-1] != "A♣" {
		t.Errorf("Expected last card A♣ but got %v", d[len(d)-1])
	}
	// Go doesn't know that running 3 tests just if Fail for any
}

func TestSaveToFileAndNewDeskFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
