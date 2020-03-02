package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, instead received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' but go %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected 'King of Clubs' but got %v", d[len(d)-1])
	}
}
