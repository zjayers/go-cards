package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Length Expectation
	expLen := 52

	if len(d) != expLen {
		t.Errorf("Expected deck length of %v, but got: %v", expLen, len(d))
	}

	// Card expectations
	expFirstCard := "Ace Of Spades"
	expLastCard := "King Of Clubs"

	// Test of the first card in the deck equals the expected card
	if d[0] != expFirstCard {
		t.Errorf("Expected first card to be: %s - but got: %s", expFirstCard, d[0])
	}

	// Test if the final card in the deck equals the expected card
	if d[len(d) - 1] != expLastCard {
		t.Errorf("Expected last card to be: %s - but got: %s", expLastCard, d[len(d) - 1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fileName := "_deck-testing"
	_ = os.Remove(fileName)

	d := newDeck()
	_ = d.saveToFile(fileName)

	deckFromFile := newDeckFromFile(fileName)

	// Length Expectation
	expLen := 52

	if len(deckFromFile) != expLen {
		t.Errorf("Expected deck length of %v, but got: %v", expLen, len(d))
	}

	_ = os.Remove(fileName)
}


