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
