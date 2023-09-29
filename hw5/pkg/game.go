package game

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/term"
)

// Raw input keycodes
const (
	up     byte = 65
	down   byte = 66
	left   byte = 68
	right  byte = 67
	escape byte = 27
	enter  byte = 13
)

type Game struct {
	Title   string
	Players []Player
	Board   Board
}

func NewGame() *Game {
	return &Game{
		Title: "TicTacToe, 'X' starts the game.\n" +
			"Press Enter to make turn.\n" +
			"Press ESC to exit.",
		Players: make([]Player, 0),
	}
}

func (game *Game) ConfigureGame() {
	for index, symbol := range Players() {
		game.Players = append(game.Players, *NewPlayer(symbol))
		if index == 0 {
			game.Players[0].ChangeTurn()
		}
	}
	game.Board = *NewBoard(3)
}

func winnerCheck(board *Board, symbol string) {
	win := ""
	draw := true
	for i := 0; i < len(board.Cells); i++ {
		win += symbol
	}

	// raw + draw
	for i := 0; i < len(board.Cells); i++ {
		result := ""
		for k := 0; k < len(board.Cells); k++ {
			result += board.Cells[i][k].Value
		}
		if strings.Contains(result, "_") {
			draw = false
		}
		if result == win {
			fmt.Printf("%s win!", symbol)
		}
	}

	if draw {
		fmt.Println("Draw!")
	}

	// col
	for i := 0; i < len(board.Cells); i++ {
		result := ""
		for k := 0; k < len(board.Cells); k++ {
			result += board.Cells[k][i].Value
		}
		if result == win {
			fmt.Printf("%s win!", symbol)
		}
	}

	// 1 diag
	result := ""
	for i := 0; i < len(board.Cells); i++ {
		for k := 0; k < len(board.Cells); k++ {
			if i == k {
				result += board.Cells[i][k].Value
			}

		}
		if result == win {
			fmt.Printf("%s win!", symbol)
			os.Exit(0)
		}
	}

	// 2 diag
	result = ""
	for i := 0; i < len(board.Cells); i++ {
		for k := 0; k < len(board.Cells); k++ {
			if i+k == len(board.Cells)-1 {
				result += board.Cells[i][k].Value
			}

		}
		if result == win {
			fmt.Printf("%s win!", symbol)
		}
	}
}

func PlayGame() {
	game := NewGame()
	game.ConfigureGame()
	fmt.Println(game.Title)
	game.Board.renderBoard(false, "X")

	for {
		for i := 0; i < len(game.Players); i++ {
			if game.Players[i].Turn {
				game.Board.makeTurn(game.Players[i].Symbol)
				game.Players[i].ChangeTurn()
				game.Players[(i+len(game.Players)-1)%len(game.Players)].ChangeTurn()
				winnerCheck(&game.Board, game.Players[i].Symbol)

			}

		}
	}

}

func getInput() byte {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		fmt.Println(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	t.Restore()
	t.Close()

	if read == 3 {
		return readBytes[2]
	} else {
		return readBytes[0]
	}
}
