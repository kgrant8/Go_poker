package main

import (
	"fmt"
	"poker/game"
)

func main() {

	var Board map[string]game.Card
	Board = make(map[string]game.Card)

	fmt.Println("welcome to my game")

	// setting up game deck and table
	table := game.NewTable(1, 2, 0, 9)
	gameDeck := game.NewDeck()

	fmt.Println("this is the table")

	player1 := game.NewPlayer("Keith", 1200)
	player2 := game.NewPlayer("Alex", 1300)

	table = game.SeatPlayer(table, player1, 0)
	table = game.SeatPlayer(table, player2, 1)

	table, gameDeck = game.Deal(table, gameDeck)
	Board = game.Flop(gameDeck, Board)

	fmt.Println(Board)

	fmt.Println("\nPLayer: ", player1)

	fmt.Println("\n", table.Seats[0], table.Seats[1])

}
