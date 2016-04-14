package game

type Table_limits struct {
	big_blind         int
	small_blind       int
	ante              int
	number_of_players int
}

type Table struct {
	Seats       map[int]Player
	TableLimits Table_limits
	Pot         int
}

func NewTable(small int, big int, ante int, numberOfPlayers int) Table {

	var tableLimits = Table_limits{big, small, ante, numberOfPlayers}
	var seats map[int]Player
	seats = make(map[int]Player)

	holdCard := Card{1, "test", "test"}

	testPlayer := Player{"test", 231, holdCard, holdCard, true, HIGH_CARD}

	for i := 0; i < numberOfPlayers; i++ {
		seats[i] = testPlayer

	}

	table := Table{seats, tableLimits, 0}

	return table

}

func SeatPlayer(table Table, player Player, seatNumber int) Table {

	table.Seats[seatNumber] = player

	return table
}
