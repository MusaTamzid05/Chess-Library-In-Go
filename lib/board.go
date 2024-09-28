package lib

import (
    "strconv"
    "fmt"
)


var BOARD = [8][8]string{
		{"♖", "♘", "♗", "♕", "♔", "♗", "♘", "♖"}, // White pieces
		{"♙", "♙", "♙", "♙", "♙", "♙", "♙", "♙"},
		{" ", " ", " ", " ", " ", " ", " ", " "}, // Empty squares
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{"♟︎", "♟︎", "♟︎", "♟︎", "♟︎", "♟︎", "♟︎", "♟︎"}, // Black pieces
		{"♜", "♞", "♝", "♛", "♚", "♝", "♞", "♜"},
	}

var CHARS string = "abcdefgh"

type Board struct {
    Cells [8][8]Cell

}

func (b Board) Render() {
    for i := 0; i < 8; i += 1 {
        for j := 0; j < 8; j += 1 {
            b.Cells[i][j].Render()
        }
        fmt.Println()
    }



}

func MakeBoard() Board {
    cells := [8][8]Cell{}

    for i := 0; i < 8; i += 1 {
        for j := 0; j < 8; j += 1 {
            name := string(CHARS[i]) + strconv.Itoa(j + 1)
            symbol := BOARD[i][j]
            cells[i][j] = MakeCell(name, symbol, i, j)
        }

    }


    board := Board{Cells: cells}
    return board

}
