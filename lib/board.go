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
    fmt.Print("  ")
    for i := 0; i < 8; i += 1 {
        fmt.Printf("%c ", CHARS[i])
    }

    fmt.Println()


    for i := 0; i < 8; i += 1 {
        rowStr := strconv.Itoa(8 - i)
        fmt.Print(rowStr, " ")
        for j := 0; j < 8; j += 1 {
            b.Cells[i][j].Render()
        }
        fmt.Println()
    }


    fmt.Print("  ")
    for i := 0; i < 8; i += 1 {
        fmt.Printf("%c ", CHARS[i])
    }

    fmt.Println()

}

func (b Board) ShowInfo() {
    for i := 0; i < 8; i += 1 {
        for j := 0; j < 8; j += 1 {
            fmt.Print(b.Cells[i][j].InfoString(),"\t")
        }
        fmt.Println()
    }


}

func (b* Board) UpdateBoard(fromCellName, toCellName string) (Cell, bool, error) {
    fromRow, fromCol, err := ConvertCellNameToRowCol(fromCellName)

    if err != nil {
        return Cell{}, false, err
    }

    symbol := b.Cells[fromRow][fromCol].Symbol
    b.Cells[fromRow][fromCol].Reset()

    toRow, toCol, err := ConvertCellNameToRowCol(toCellName)

    if err != nil {
        return Cell{}, false, err
    }

    pieceRemoved := b.Cells[toRow][toCol].IsEmpty() == false
    b.Cells[toRow][toCol].Symbol = symbol

    return b.Cells[toRow][toCol], pieceRemoved, nil




}



func MakeBoard() Board {
    cells := [8][8]Cell{}

    for i := 0; i < 8; i += 1 {
        rowStr := strconv.Itoa(8 - i)
        for j := 0; j < 8; j += 1 {
            name := string(CHARS[j]) + rowStr
            symbol := BOARD[i][j]
            cells[i][j] = MakeCell(name, symbol, i, j)
        }

    }


    board := Board{Cells: cells}
    return board

}
