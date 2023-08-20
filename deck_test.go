package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length to be 16, but got", len(d))
	}

	if d[0] != "Ace Of Spades" {
		t.Error("Expected first card in deck to be Ace Of Spades, but it is", d[0])
	}

	if d[len(d)-1] != "King Of Clubs" {
		t.Error("Expected first card in deck to be King Of Clubs, but it is", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testdeck")

	d := newDeck()
	d.saveFile("_testdeck")

	deckFromFile, err := newDeckFromFile("_testdeck")

	if err != nil {
		t.Error("Error while testing", err)
	}

	if len(deckFromFile) != 16 {
		t.Error("Expected deck length to be 16, but got", len(deckFromFile))
	}

	if deckFromFile[0] != "Ace Of Spades" {
		t.Error("Expected first card in deck to be Ace Of Spades, but it is", deckFromFile[0])
	}

	if deckFromFile[len(d)-1] != "King Of Clubs" {
		t.Error("Expected first card in deck to be King Of Clubs, but it is", deckFromFile[len(d)-1])
	}

	os.Remove("_testdeck")
}
