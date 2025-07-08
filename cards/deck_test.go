package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 2000 {
		t.Errorf("sth wrong %v", len(cards))
	}
}
