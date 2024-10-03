package lib

import (
    "strconv"
)

func ConvertCellNameToRowCol(cellName string) (int, int, error) {
    var row int
    var col int

    for index, char := range CHARS {
        if byte(char)== cellName[0] {
            col = index
            break

        }
    }

    row, err := strconv.Atoi(string(cellName[1]))

    if err != nil {
        return row, col, err
    }

    return 8 - row, col, nil
}

type Move struct {
    Row int
    Col int
}


func (m Move) GetString() string {
    return string(CHARS[m.Col]) + strconv.Itoa(8 - m.Row)
}


func (m Move) Equal(other Move) bool{
    return m.Row == other.Row && m.Col == other.Col
}


func MakeMove(cellName string) (Move, error){
    row, col, err := ConvertCellNameToRowCol(cellName)
    return Move{Row: row, Col: col}, err
}
