package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Error: Yooo How the hell did you not manage to get a full deck of card?!?!, You somehow managed to get %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Error: Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Error: Expected King of Hearts but got %v", d[len(d)-1])
	}
}
