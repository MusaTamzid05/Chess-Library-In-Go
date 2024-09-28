package lib

import (
    "fmt"
)

type Cell struct {
    Name string
    Symbol string

    Row int
    Col int

}

var EMPTYCELL string = "."

func (cell  Cell) Render() {
    fmt.Print(cell.Symbol, " ")
}

func (cell  Cell) InfoString() string  {
    return fmt.Sprintf("[%d][%d] %s %s", 
    cell.Row,
    cell.Col,
    cell.Name,
    cell.Symbol,
   )
}

func MakeCell(name, symbol string, row, col int) Cell {

    if symbol == " " {
        symbol = EMPTYCELL
    }

    return Cell{
        Name: name,
        Symbol: symbol,
        Row: row,
        Col: col,
    }
}
