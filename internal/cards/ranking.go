package cards

import (
	"cmp"
	"errors"
	"slices"

	"github.com/bobby-palmer/degen-now/internal/snacks"
)

type HandRank int

const (
	NilRank HandRank = iota // specical base case
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type HandResult struct {
	Rank        HandRank
	Tiebreakers []Rank
}

func (h HandResult) Compare(other HandResult) int {
	if h.Rank != other.Rank {
		return cmp.Compare(h.Rank, other.Rank)
	}

	return slices.Compare(h.Tiebreakers, other.Tiebreakers)
}

var ErrNotEnoughCards = errors.New("expected atleast 5 cards in hand")

func Rank5(hand []Card) (HandResult, error) {

	if len(hand) < 5 {
		return HandResult{}, ErrNotEnoughCards
	}

	if result := getRoyalFlush(hand); result != nil {
		return *result, nil
	}

	if result := getStraightFlush(hand); result != nil {
		return *result, nil
	}

	if result := getQuads(hand); result != nil {
		return *result, nil
	}

	if result := getFullHouse(hand); result != nil {
		return *result, nil
	}

	if result := getFlush(hand); result != nil {
		return *result, nil
	}

	if result := getStraight(hand); result != nil {
		return *result, nil
	}

	if result := getTrips(hand); result != nil {
		return *result, nil
	}

	if result := getTwoPair(hand); result != nil {
		return *result, nil
	}

	if result := getOnePair(hand); result != nil {
		return *result, nil
	}

	return getHighCard(hand), nil
}

func getRoyalFlush(hand []Card) *HandResult {
	for _, suit := range Suits {
		if slices.Contains(hand, Card{suit, RankAce}) &&
			slices.Contains(hand, Card{suit, RankKing}) &&
			slices.Contains(hand, Card{suit, RankQueen}) &&
			slices.Contains(hand, Card{suit, RankJack}) &&
			slices.Contains(hand, Card{suit, RankTen}) {
			return &HandResult{RoyalFlush, nil}
		}
	}

	return nil
}

func getStraightCandidates(starting Rank) []Rank {
	if starting == RankAce {
		return []Rank{RankAce, RankTwo, RankThree, RankFour, RankFive}
	} else if starting <= RankTen {
		return []Rank{
			starting,
			starting + 1,
			starting + 2,
			starting + 3,
			starting + 4,
		}
	}

	return nil
}

func getStraightFlush(hand []Card) *HandResult {

	result := HandResult{}

	for _, suit := range Suits {
		for _, rank := range Ranks {

			candidates := getStraightCandidates(rank)
			if candidates == nil {
				continue
			}

			if !snacks.AllOf(candidates, func(r Rank) bool {
				return slices.Contains(hand, Card{suit, r})
			}) {
				continue
			}

			slices.Reverse(candidates)

			thisResult := HandResult{StraightFlush, candidates}
			if result.Compare(thisResult) < 0 {
				result = thisResult
			}

		}
	}

	if result.Rank != StraightFlush {
		return nil
	}

	return &result
}

func getQuads(hand []Card) *HandResult {

	result := HandResult{}

	for _, rank := range Ranks {

		if !snacks.AllOf(Suits, func(s Suit) bool {
			return slices.Contains(hand, Card{s, rank})
		}) {
			continue
		}

		notQuads := snacks.Filter(hand, func(c Card) bool {
			return c.Rank != rank
		})

		kicker := slices.MaxFunc(notQuads, func(a, b Card) int {
			return cmp.Compare(a.Rank, b.Rank)
		})

		thisResult := HandResult{FourOfAKind, []Rank{rank, kicker.Rank}}
		if result.Compare(thisResult) < 0 {
			result = thisResult
		}

	}

	if result.Rank != FourOfAKind {
		return nil
	}

	return &result
}

func getFullHouse(hand []Card) *HandResult {
	result := HandResult{}

	for _, three := range Ranks {
		for _, two := range Ranks {

			if three == two {
				continue
			}

			if snacks.Count(hand, func(c Card) bool {
				return c.Rank == three
			}) < 3 {
				continue
			}

			if snacks.Count(hand, func(c Card) bool {
				return c.Rank == two
			}) < 2 {
				continue
			}

			thisResult := HandResult{FullHouse, []Rank{three, two}}
			if result.Compare(thisResult) < 0 {
				result = thisResult
			}
		}
	}

	if result.Rank != FullHouse {
		return nil
	}

	return &result
}

func getFlush(hand []Card) *HandResult {
	result := HandResult{}

	for _, suit := range Suits {

		cardsOfSuit := snacks.Filter(hand, func(c Card) bool {
			return c.Suit == suit
		})

		if len(cardsOfSuit) < 5 {
			continue
		}

		ranksOfSuit := snacks.Map(cardsOfSuit, func(c Card) Rank {
			return c.Rank
		})

		slices.Sort(ranksOfSuit)
		slices.Reverse(ranksOfSuit)

		thisResult := HandResult{Flush, ranksOfSuit[:5]}
		if result.Compare(thisResult) < 0 {
			result = thisResult
		}

	}

	if result.Rank != Flush {
		return nil
	}

	return &result
}

func getStraight(hand []Card) *HandResult {
	result := HandResult{}

	ranks := snacks.Map(hand, func(c Card) Rank { return c.Rank })

	for _, rank := range Ranks {

		candidates := getStraightCandidates(rank)
		if candidates == nil {
			continue
		}

		if !snacks.AllOf(candidates, func(r Rank) bool {
			return slices.Contains(ranks, r)
		}) {
			continue
		}

		slices.Reverse(candidates)

		thisResult := HandResult{Straight, candidates}
		if result.Compare(thisResult) < 0 {
			result = thisResult
		}
	}

	if result.Rank != Straight {
		return nil
	}

	return &result
}

func getTrips(hand []Card) *HandResult {
	result := HandResult{}

	ranks := snacks.Map(hand, func(c Card) Rank {
		return c.Rank
	})
	slices.Sort(ranks)
	slices.Reverse(ranks)

	for _, rank := range Ranks {

		if snacks.Count(hand, func(c Card) bool {
			return c.Rank == rank
		}) < 3 {
			continue
		}

		otherRanks := snacks.Filter(ranks, func(r Rank) bool {
			return r != rank
		})

		thisResult := HandResult{ThreeOfAKind, otherRanks[:2]}
		if result.Compare(thisResult) < 0 {
			result = thisResult
		}

	}

	if result.Rank != ThreeOfAKind {
		return nil
	}

	return &result
}

func getTwoPair(hand []Card) *HandResult {

	result := HandResult{}

	ranks := snacks.Map(hand, func(c Card) Rank {
		return c.Rank
	})
	slices.Sort(ranks)
	slices.Reverse(ranks)

	for _, pairOne := range Ranks {
		for _, pairTwo := range Ranks {

			if pairOne == pairTwo {
				continue
			}

			if snacks.Count(hand, func(c Card) bool { return c.Rank == pairOne }) < 2 {
				continue
			}

			if snacks.Count(hand, func(c Card) bool { return c.Rank == pairTwo }) < 2 {
				continue
			}

			withoutPairOne := snacks.Filter(ranks, func(r Rank) bool {
				return r != pairOne
			})

			withoutPairs := snacks.Filter(withoutPairOne, func(r Rank) bool {
				return r != pairTwo
			})

			thisResult := HandResult{TwoPair, withoutPairs[:1]}
			if result.Compare(thisResult) < 0 {
				result = thisResult
			}

		}
	}

	if result.Rank != TwoPair {
		return nil
	}

	return &result
}

func getOnePair(hand []Card) *HandResult {

	result := HandResult{}

	ranks := snacks.Map(hand, func(c Card) Rank { return c.Rank })
	slices.Sort(ranks)
	slices.Reverse(ranks)

	for _, rank := range Ranks {

		if snacks.Count(ranks, func(r Rank) bool { return r == rank }) < 2 {
			continue
		}

		withoutPair := snacks.Filter(ranks, func(r Rank) bool {
			return r != rank
		})

		thisResult := HandResult{OnePair, withoutPair[:3]}
		if result.Compare(thisResult) < 0 {
			result = thisResult
		}

	}

	if result.Rank != OnePair {
		return nil
	}

	return &result
}

func getHighCard(hand []Card) HandResult {

	ranks := snacks.Map(hand, func(c Card) Rank { return c.Rank })
	slices.Sort(ranks)
	slices.Reverse(ranks)

	return HandResult{HighCard, ranks[:5]}
}
