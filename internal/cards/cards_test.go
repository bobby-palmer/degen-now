package cards

import (
	"slices"
	"testing"
)

func TestDeckSize(t *testing.T) {
	deck := NewDeck()

	if len(deck) != 52 {
		t.Errorf("expected 52 cards, got: %d", len(deck))
	}
}

func TestDeckShuffle(t *testing.T) {
	deck := NewDeck()

	shuffled := slices.Clone(deck)

	shuffled.Shuffle()

	if slices.Equal(deck, shuffled) {
		t.Errorf("expected deck to differ after shuffling")
	}
}

func TestDeckDeal(t *testing.T) {
	deck := NewDeck()

	for i := range 52 {
		_, err := deck.Deal()
		if err != nil {
			t.Errorf("failed to deal card: %d", i+1)
		}
	}

	var err error
	_, err = deck.Deal()
	if err == nil {
		t.Errorf("expected deck to be empty")
	}
}
