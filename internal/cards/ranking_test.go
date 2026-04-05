package cards

import (
	"testing"
)

func TestRank5(t *testing.T) {
	type TestPair struct {
		hand     []Card
		expected HandRank
	}

	toTest := []TestPair{
		// Royal flush
		{[]Card{{SuitSpades, RankAce}, {SuitSpades, RankKing}, {SuitSpades, RankQueen}, {SuitSpades, RankJack}, {SuitSpades, RankTen}}, RoyalFlush},
		{[]Card{{SuitClubs, RankAce}, {SuitClubs, RankKing}, {SuitClubs, RankQueen}, {SuitClubs, RankJack}, {SuitClubs, RankTen}}, RoyalFlush},
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankKing}, {SuitHearts, RankQueen}, {SuitHearts, RankJack}, {SuitHearts, RankTen}}, RoyalFlush},
		{[]Card{{SuitDiamonds, RankAce}, {SuitDiamonds, RankKing}, {SuitDiamonds, RankQueen}, {SuitDiamonds, RankJack}, {SuitDiamonds, RankTen}}, RoyalFlush},
		// Straight Flush
		{[]Card{{SuitSpades, RankKing}, {SuitSpades, RankQueen}, {SuitSpades, RankJack}, {SuitSpades, RankTen}, {SuitSpades, RankNine}}, StraightFlush},
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankTwo}, {SuitHearts, RankThree}, {SuitHearts, RankFour}, {SuitHearts, RankFive}}, StraightFlush},
	}

	for _, pair := range toTest {

		result, err := Rank5(pair.hand)
		if err != nil {
			t.Errorf("getting hand rank: %v", err)
			continue
		}

		if result.Rank != pair.expected {
			t.Errorf("got %v wanted %v", result.Rank, pair.expected)
			continue
		}
	}
}
