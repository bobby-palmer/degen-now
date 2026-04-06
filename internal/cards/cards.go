package cards

import (
	"errors"
	"math/rand"
)

type Suit int

const (
	SuitHearts Suit = iota
	SuitDiamonds
	SuitSpades
	SuitClubs
)

func (s Suit) String() string {
	return [...]string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}[s]
}

var Suits = []Suit{
	SuitHearts,
	SuitDiamonds,
	SuitSpades,
	SuitClubs,
}

type Rank int

const (
	RankTwo Rank = iota
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

func (r Rank) String() string {
	return [...]string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}[r]
}

var Ranks = []Rank{
	RankTwo,
	RankThree,
	RankFour,
	RankFive,
	RankSix,
	RankSeven,
	RankEight,
	RankNine,
	RankTen,
	RankJack,
	RankQueen,
	RankKing,
	RankAce,
}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	return c.Rank.String() + " of " + c.Suit.String()
}

type Deck []Card

func NewDeck() Deck {
	deck := make([]Card, 0, 52)

	for _, suit := range Suits {
		for _, rank := range Ranks {
			deck = append(deck, Card{suit, rank})
		}
	}

	return deck
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

var ErrNoCards = errors.New("no cards to deal")

func (d *Deck) Deal() (Card, error) {
	if len(*d) == 0 {
		return Card{}, ErrNoCards
	}

	result := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]

	return result, nil
}
