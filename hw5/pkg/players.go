package game

type Player struct {
	Symbol string
	Turn   bool
}

func Players() map[int]string {
	return map[int]string{
		0: "X",
		1: "O",
	}
}

func NewPlayer(symbol string) *Player {
	return &Player{
		Symbol: symbol,
	}
}

func (player *Player) ChangeTurn() {
	if player.Turn == true {
		player.Turn = false
	} else {
		player.Turn = true
	}
}
