package game

type Player struct {
	name     string
	balance  float32
	CardOne  Card
	CardTwo  Card
	folded   bool
	BestHand PokerHand
}

func NewPlayer(name string, amount float32) Player {

	holdCard := Card{1, "test", "test"}

	var newPlayer = Player{name, amount, holdCard, holdCard, true, HIGH_CARD}

	return newPlayer

}
