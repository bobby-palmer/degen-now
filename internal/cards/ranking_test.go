package cards

import (
	"testing"
)

func TestRank5CorrectRank(t *testing.T) {
	type TestPair struct {
		hand     []Card
		expected HandRank
	}

	// TODO add names to harness
	toTest := []TestPair{
		// Royal flush
		{[]Card{{SuitSpades, RankAce}, {SuitSpades, RankKing}, {SuitSpades, RankQueen}, {SuitSpades, RankJack}, {SuitSpades, RankTen}}, RoyalFlush},
		{[]Card{{SuitClubs, RankAce}, {SuitClubs, RankKing}, {SuitClubs, RankQueen}, {SuitClubs, RankJack}, {SuitClubs, RankTen}}, RoyalFlush},
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankKing}, {SuitHearts, RankQueen}, {SuitHearts, RankJack}, {SuitHearts, RankTen}}, RoyalFlush},
		{[]Card{{SuitDiamonds, RankAce}, {SuitDiamonds, RankKing}, {SuitDiamonds, RankQueen}, {SuitDiamonds, RankJack}, {SuitDiamonds, RankTen}}, RoyalFlush},
		// Straight Flush
		{[]Card{{SuitSpades, RankKing}, {SuitSpades, RankQueen}, {SuitSpades, RankJack}, {SuitSpades, RankTen}, {SuitSpades, RankNine}}, StraightFlush},
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankTwo}, {SuitHearts, RankThree}, {SuitHearts, RankFour}, {SuitHearts, RankFive}}, StraightFlush},
		// Quads
		{[]Card{{SuitSpades, RankFive}, {SuitHearts, RankFive}, {SuitDiamonds, RankFive}, {SuitClubs, RankFive}, {SuitClubs, RankNine}}, FourOfAKind},
		// Full house
		{[]Card{{SuitHearts, RankTen}, {SuitDiamonds, RankTen}, {SuitSpades, RankTen}, {SuitHearts, RankFive}, {SuitClubs, RankFive}}, FullHouse},
		// Flush
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankTwo}, {SuitHearts, RankThree}, {SuitHearts, RankFour}, {SuitHearts, RankSix}}, Flush},
		// Straight
		{[]Card{{SuitHearts, RankAce}, {SuitHearts, RankTwo}, {SuitHearts, RankThree}, {SuitHearts, RankFour}, {SuitClubs, RankFive}}, Straight},
		{[]Card{{SuitSpades, RankAce}, {SuitDiamonds, RankKing}, {SuitDiamonds, RankQueen}, {SuitDiamonds, RankJack}, {SuitDiamonds, RankTen}}, Straight},
		// Trips
		{[]Card{{SuitSpades, RankTwo}, {SuitClubs, RankTwo}, {SuitHearts, RankTwo}, {SuitHearts, RankFive}, {SuitHearts, RankTen}}, ThreeOfAKind},
		// two pair
		{[]Card{{SuitSpades, RankFour}, {SuitClubs, RankFour}, {SuitSpades, RankFive}, {SuitClubs, RankFive}, {SuitHearts, RankNine}}, TwoPair},
		// one pair
		{[]Card{{SuitSpades, RankAce}, {SuitHearts, RankAce}, {SuitClubs, RankNine}, {SuitClubs, RankEight}, {SuitClubs, RankSeven}}, OnePair},
		// high card
		{[]Card{{SuitClubs, RankAce}, {SuitSpades, RankThree}, {SuitClubs, RankFour}, {SuitClubs, RankFive}, {SuitClubs, RankSeven}}, HighCard},
		{[]Card{{SuitClubs, RankKing}, {SuitClubs, RankAce}, {SuitHearts, RankTwo}, {SuitClubs, RankThree}, {SuitClubs, RankFour}}, HighCard},
	}

	for i, pair := range toTest {

		result, err := Rank5(pair.hand)
		if err != nil {
			t.Errorf("getting hand rank: %v", err)
			continue
		}

		if result.Rank != pair.expected {
			t.Errorf("%d: got %v wanted %v", i, result.Rank, pair.expected)
			continue
		}
	}
}

func TestRank5Tiebreaker(t *testing.T) {
	// TODO test ranking within rank
}
