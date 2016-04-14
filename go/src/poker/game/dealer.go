package game

import (
	"fmt"
)

type PokerHand int

const (
	HIGH_CARD PokerHand = 1 + iota
	PAIR
	TWO_PAIR
	THREE_OFAKIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OFAKIND
	STRAIGHT_FLUSH
)

// //var Board map[string]Card

// func HoldCards(deck []Card, table Table_limits) []Card {
// 	// c = top card in the deck
// 	// d = new deck after top card been removed

// 	for i := 0; i <= table.number_of_players*2; i++ {
// 		//c := deck[0]
// 		deck = append(deck[:0], deck[0+1:]...)
// 	}

// 	return deck

// }

func Flop(deck []Card, board map[string]Card) map[string]Card {

	fmt.Println("the deck: ")
	fmt.Println(deck[0], deck[1], deck[2])

	board["flop1"] = deck[1]
	board["flop2"] = deck[2]
	board["flop3"] = deck[3]
	board["turn"] = deck[5]
	board["river"] = deck[7]

	deck = append(deck[:0], deck[0+7:]...)

	return board

}

func Deal(table Table, deck []Card) (Table, []Card) {

	for j := 0; j < 2; j++ {

		for i := 0; i < 2; i++ {

			if j == 0 {
				temp := table.Seats[i]
				temp.CardOne = deck[0]
				table.Seats[i] = temp

			} else {
				temp := table.Seats[i]
				temp.CardTwo = deck[0]
				table.Seats[i] = temp

			}

			deck = append(deck[:0], deck[0+1:]...)

		}
	}

	return table, deck

}

/*func WinningHand(table Table, board map[string]Card) Player {

	for i := 0; i < table.TableLimits.number_of_players; i++ {

		if

	}
}
*/
func PlayersBestHand(player Player, board map[string]Card) {

	flushPossible := false
	fullHousePossible := false

	for i := 0; i < 5; i++ {
		flushSuit := "suit"
		num := 0

		switch i {
		case 1:
			flushSuit = "Spades"
			num, flushSuit = FlushCheck(flushSuit, board)
		case 2:
			flushSuit = "Hearts"
			num, flushSuit = FlushCheck(flushSuit, board)
		case 3:
			flushSuit = "Clubs"
			num, flushSuit = FlushCheck(flushSuit, board)
		case 4:
			flushSuit = "Diamonds"
			num, flushSuit = FlushCheck(flushSuit, board)

		}
		if num >= 3 {
			flushPossible = true
		}
	}

	if player.folded {

	}
	flushPossible = false
}

func FlushCheck(testSuit string, board map[string]Card) (string, bool) {
	num := 0
	possible := false

	if board["flop1"].Suit == testSuit {
		num++
	}

	if board["flop2"].Suit == testSuit {
		num++
	}

	if board["flop3"].Suit == testSuit {
		num++
	}

	if board["turn"].Suit == testSuit {
		num++
	}

	if board["river"].Suit == testSuit {
		num++
	}
	if num >= 3 {
		possible = true
	}
	return testSuit, possible

}
