package game

import (
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	number int
	Suit   string
	Alas   string
}

func GetRandomNumber(min, max int) int {

	return rand.Intn(max-min) + min
}

func Shuffle(shuffleDeck []Card) []Card {

	rand.Seed(time.Now().Unix())

	for i := 0; i < 200; i++ {
		j := GetRandomNumber(0, 52)
		k := GetRandomNumber(0, 52)

		shuffleDeck[j], shuffleDeck[k] = shuffleDeck[k], shuffleDeck[j]

	}

	return shuffleDeck
}

func NewDeck() []Card {

	var deck = make([]Card, 0)
	card := Card{1, "hearts", "Ace"}
	for i := 0; i < 52; i++ {

		card.number = (i % 13) + 1

		switch k := (i % 13) + 1; k {
		case 1:
			card.Alas = "Ace"
		case 11:
			card.Alas = "Jack"
		case 12:
			card.Alas = "Queen"
		case 13:
			card.Alas = "King"
		default:
			card.Alas = strconv.Itoa(k)

		}

		switch {
		case 0 <= i && i < 13:
			card.Suit = "Spades"
		case 13 <= i && i < 26:
			card.Suit = "Clubs"
		case 26 <= i && i < 39:
			card.Suit = "Hearts"
		case 39 <= i && i < 53:
			card.Suit = "Diamonds"
		}

		deck = append(deck, card)

	}
	deck = Shuffle(deck)

	return deck

}
