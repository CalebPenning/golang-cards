package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 60 {
		t.Errorf("Expected deck length of 60 but got %v", len(d))
	}
}