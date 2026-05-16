package game

import (
	"fmt"

	"github.com/bobby-palmer/degen-now/internal/cards"
)

type Config struct {
	BombPot int64
}

// Return a default table config
func NewConfig() *Config {
	return &Config{
		BombPot: 1,
	}
}

type Player struct {
	Stack      int64
	CurrentBet int64
	PutInPot   int64
	Cards      []cards.Card
	IsInHand   bool
}

type Street int

const (
	StreetNone Street = iota
	StreetDiscard
	StreetFlop
	StreetTurn
	StreetRiver
	StreetShowdown
)

type Table struct {
	Config  *Config
	Players []Player
	Deck    cards.Deck
	Board   []cards.Card
	Street  Street
	Dealer  int
	Action  int
}

func NewTable() *Table {
	return &Table{
		Config: NewConfig(),
	}
}

func (t *Table) AddPlayer(stack int64) (int, error) {
	if t.Street != StreetNone {
		return 0, fmt.Errorf("cannot add player mid game")
	}

	// TODO max player logic

	t.Players = append(t.Players, Player{Stack: stack})

	return len(t.Players) - 1, nil
}

func (t *Table) Next() error {

	if t.Action != -1 {
		return fmt.Errorf("street not finished")
	}

	switch t.Street {
	case StreetNone:
		t.Deck = cards.NewDeck()
		t.Deck.Shuffle()

		for i := range t.Players {
			t.Players[i].IsInHand = true
			t.Players[i].PutInPot = t.Config.BombPot
			t.Players[i].Stack -= t.Config.BombPot
			t.Players[i].Cards = []cards.Card{}

			for _ = range 5 {
				c, err := t.Deck.Deal()
				if err != nil {
					return fmt.Errorf("dealing player %d: %w", i, err)
				}

				t.Players[i].Cards = append(t.Players[i].Cards, c)
			}
		}

		t.Board = []cards.Card{}

		for _ = range 3 {
			c, err := t.Deck.Deal()
			if err != nil {
				return fmt.Errorf("dealing board: %w", err)
			}

			t.Board = append(t.Board, c)
		}

		t.Action = t.Dealer
		t.Street = StreetDiscard

	case StreetDiscard:
		t.Action = t.Dealer
		t.Street = StreetFlop

	case StreetFlop:
		c, err := t.Deck.Deal()
		if err != nil {
			return fmt.Errorf("dealing turn: %w", err)
		}

		t.Board = append(t.Board, c)
		t.Action = t.Dealer
		t.Street = StreetTurn

	case StreetTurn:
		c, err := t.Deck.Deal()
		if err != nil {
			return fmt.Errorf("dealing river: %w", err)
		}

		t.Board = append(t.Board, c)
		t.Action = t.Dealer
		t.Street = StreetRiver

	case StreetRiver:
		// TODO scoring and pot distribution

	}

	return nil
}

func (t *Table) Discard(playerID int, cardID int) error

func (t *Table) Check(playerID int) error

func (t *Table) Bet(playerID int, amount int) error

func (t *Table) Fold(playerID int) error
