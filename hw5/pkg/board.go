package game

import (
	"fmt"
	"os"
)

type Board struct {
	CursorPosX int
	CursorPosY int
	Cells      [][]Cell
}

type Cell struct {
	Value string
	ID    int
}

func (cell *Cell) setValue(value string) {
	cell.Value = value
}

func NewBoard(size int) *Board {
	board := &Board{
		Cells: make([][]Cell, size),
	}
	for i := 0; i < size; i++ {
		board.Cells[i] = make([]Cell, size)
		for k := 0; k < size; k++ {
			board.Cells[i][k].Value = "_"
		}
	}
	return board
}

func (board *Board) renderBoard(redraw bool, symbol string) {
	if redraw {
		fmt.Printf("\033[%dA", len(board.Cells))
	}

	for i, _ := range board.Cells {

		for k, cell := range board.Cells[i] {
			var printCell string
			delimeter := " |"
			if k == len(board.Cells[k])-1 {
				delimeter = "\n"
			}

			if cell.Value != "_" && i == board.CursorPosX && k == board.CursorPosY {
				printCell = ">" + cell.Value
			} else if cell.Value == "_" && i == board.CursorPosX && k == board.CursorPosY {
				printCell = ">" + symbol
			} else {
				printCell = " " + cell.Value
			}

			fmt.Printf(" %s%s", printCell, delimeter)
		}
	}
}

func (board *Board) makeTurn(symbol string) {

	for {
		keyCode := getInput()
		if keyCode == escape {
			os.Exit(0)
		} else if keyCode == enter {
			if board.Cells[board.CursorPosX][board.CursorPosY].Value == "_" {
				board.Cells[board.CursorPosX][board.CursorPosY].setValue(symbol)
				return
			}
		} else if keyCode == up {
			board.CursorPosX = (board.CursorPosX + len(board.Cells) - 1) % len(board.Cells)
			board.renderBoard(true, symbol)
		} else if keyCode == down {
			board.CursorPosX = (board.CursorPosX + 1) % len(board.Cells)
			board.renderBoard(true, symbol)
		} else if keyCode == left {
			board.CursorPosY = (board.CursorPosY + len(board.Cells[0]) - 1) % len(board.Cells[0])
			board.renderBoard(true, symbol)
		} else if keyCode == right {
			board.CursorPosY = (board.CursorPosY + 1) % len(board.Cells[0])
			board.renderBoard(true, symbol)
		}
	}
}
