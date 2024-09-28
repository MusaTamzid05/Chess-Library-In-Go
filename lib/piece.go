package lib

type RootPiece struct {
    Cell Cell
}


type Piece interface {
    GetCommon() *RootPiece
    GetValidMoves() []Move
}

type Rook struct {
    Common RootPiece
}

func MakeRook(cell Cell) Rook {
    return Rook{Common: RootPiece{Cell: cell}}
}

func (p Rook) GetCommon() *RootPiece {
    return &p.Common
}


func (p Rook) GetValidMoves() []Move {
    moves := []Move{}

    row := p.Common.Cell.Row
    col := p.Common.Cell.Col

    newRow:= row
    newCol := col

    for i := col + 1; i < 8; i += 1 {
        newCol = i 
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    for i := col - 1; i >= 0; i -= 1 {
        newCol = i 
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    newRow = row
    newCol = col


    for i := row + 1; i < 8; i += 1 {
        newRow = i 
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    for i := row - 1; i >= 0; i -= 1 {
        newRow = i 
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    return moves
}


