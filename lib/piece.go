package lib

type RootPiece struct {
    Cell Cell
}


type Piece interface {
    GetCommon() *RootPiece
    GetValidMoves() []Move
}

type Rook struct {
    Common *RootPiece
}

func MakeRook(cell Cell) Rook {
    return Rook{Common: &RootPiece{Cell: cell}}
}

func (p Rook) GetCommon() *RootPiece {
    return p.Common
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


type Knight struct {
    Common *RootPiece
}

func MakeKnight(cell Cell) Knight {
    return Knight{Common: &RootPiece{Cell: cell}}
}

func (p Knight) GetCommon() *RootPiece {
    return p.Common
}


func (p Knight) GetValidMoves() []Move {
    moves := []Move{}

    row := p.Common.Cell.Row
    col := p.Common.Cell.Col

    newRow := row + 2
    newCol := col + 1


    if newRow < 8 && newCol < 8 {
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    newRow = row + 2
    newCol = col - 1


    if newRow < 8 && newCol >= 0 {
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }


    newRow = row - 2
    newCol = col - 1


    if newRow >= 0 && newCol >= 0 {
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }



    newRow = row - 2
    newCol = col + 1


    if newRow >= 0 && newCol < 8 {
        moves = append(moves, Move{Row: newRow, Col: newCol})
    }

    return moves
}


type Bishop struct {
    Common *RootPiece
}

func MakeBishop(cell Cell) Bishop {
    return Bishop{Common: &RootPiece{Cell: cell}}
}

func (p Bishop) GetCommon() *RootPiece {
    return p.Common
}


func (p Bishop) GetValidMoves() []Move {
    moves := []Move{}

    row := p.Common.Cell.Row
    col := p.Common.Cell.Col

    newRow := row
    newCol := col

    // down, right
    for true {
        newRow += 1
        newCol += 1


        if newRow < 8 && newCol < 8 {
            moves = append(moves, Move{Row: newRow, Col: newCol})
        } else {
            break
        }
    }

    newRow = row
    newCol = col


    // up, left
    for true {
        newRow -= 1
        newCol -= 1

        if newRow >= 0  && newCol >= 0  {
            moves = append(moves, Move{Row: newRow, Col: newCol})
        } else {
            break
        }
    }

    newRow = row
    newCol = col

    // down , left
    for true {
        newRow += 1
        newCol -= 1

        if newRow < 8  && newCol >= 0  {
            moves = append(moves, Move{Row: newRow, Col: newCol})
        } else {
            break
        }
    }

    newRow = row
    newCol = col

    // down , right
    for true {
        newRow -= 1
        newCol += 1

        if newRow >= 0  && newCol < 8  {
            moves = append(moves, Move{Row: newRow, Col: newCol})
        } else {
            break
        }
    }


    return moves
}


type Queen struct {
    Common *RootPiece
}

func MakeQueen(cell Cell) Queen {
    return Queen{Common: &RootPiece{Cell: cell}}
}

func (p Queen) GetCommon() *RootPiece {
    return p.Common
}


func (p Queen) GetValidMoves() []Move {
    moves := []Move{}

    rook := MakeRook(p.Common.Cell)
    moves = append(moves, rook.GetValidMoves()...)

    bishop := MakeBishop(p.Common.Cell)
    moves = append(moves, bishop.GetValidMoves()...)

    return moves
}




type King struct {
    Common *RootPiece
}

func MakeKing(cell Cell) King {
    return King{Common: &RootPiece{Cell: cell}}
}

func (p King) GetCommon() *RootPiece {
    return p.Common
}


func (p King) GetValidMoves() []Move {
    moves := []Move{}

    row := p.Common.Cell.Row
    col := p.Common.Cell.Col

    if row - 1 >= 0 {
        moves = append(moves, Move{Row: row - 1, Col: col})

    }

    if row + 1 < 8 {
        moves = append(moves, Move{Row: row + 1, Col: col})
    }


    if col - 1 >= 0 {
        moves = append(moves, Move{Row: row, Col: col - 1})

    }

    if col + 1 < 8 {
        moves = append(moves, Move{Row: row, Col: col + 1})
    }


    if row - 1 >= 0 && col - 1 >= 0 {
        moves = append(moves, Move{Row: row - 1, Col: col - 1})
 
    }

    if row + 1 < 8 && col + 1 < 8 {
        moves = append(moves, Move{Row: row + 1, Col: col + 1})
    }

    if row - 1 >= 0 && col + 1 < 8 {
        moves = append(moves, Move{Row: row - 1, Col: col + 1})
    }

    if row + 1 < 8 && col - 1 >= 0 {
        moves = append(moves, Move{Row: row + 1, Col: col - 1})
    }


    return moves
}


type Pawn struct {
    Common *RootPiece
    white bool
}

func MakePawn(cell Cell, white bool) Pawn {
    return Pawn{Common: &RootPiece{Cell: cell}, white: white}
}

func (p Pawn) GetCommon() *RootPiece {
    return p.Common
}


func (p Pawn) GetValidMoves() []Move {
    moves := []Move{}

    row := p.Common.Cell.Row
    col := p.Common.Cell.Col

    if p.white {
        if row - 1 >= 0 {
            moves = append(moves, Move{Row: row - 1, Col: col})

        }

    } else {

        if row + 1 < 8 {
            moves = append(moves, Move{Row: row +  1, Col: col})

        }

    }


    return moves
}


