package lib


func InitBlackPieces(board Board) map[string]Piece {
    pieces := map[string]Piece{}

    pieces["a8"] = MakeRook(board.Cells[0][0])
    pieces["b8"] = MakeKnight(board.Cells[0][1])
    pieces["c8"] = MakeBishop(board.Cells[0][2])
    pieces["d8"] = MakeQueen(board.Cells[0][3])
    pieces["e8"] = MakeKing(board.Cells[0][4])
    pieces["f8"] = MakeBishop(board.Cells[0][5])
    pieces["g8"] = MakeKnight(board.Cells[0][6])
    pieces["h8"] = MakeRook(board.Cells[0][7])

    for i := 0; i < 8; i += 1 {
        pieces[string(CHARS[i]) + "7"] = MakePawn(board.Cells[1][i], false)
    }

    return pieces

}


func InitWhitePieces(board Board) map[string]Piece {
    pieces := map[string]Piece{}

    for i := 0; i < 8; i += 1 {
        pieces[string(CHARS[i]) + "2"] = MakePawn(board.Cells[6][i], true)
    }

    pieces["a1"] = MakeRook(board.Cells[7][0])
    pieces["b1"] = MakeKnight(board.Cells[7][1])
    pieces["c1"] = MakeBishop(board.Cells[7][2])
    pieces["d1"] = MakeQueen(board.Cells[7][3])
    pieces["e1"] = MakeKing(board.Cells[7][4])
    pieces["f1"] = MakeBishop(board.Cells[7][5])
    pieces["g1"] = MakeKnight(board.Cells[7][6])
    pieces["h1"] = MakeRook(board.Cells[7][7])


    return pieces

}


type RootPlayer struct {
    Pieces map[string]Piece
}

func (p RootPlayer) PieceExists(cellName string ) bool {
    _, ok := p.Pieces[cellName]
    return ok
}

func (p RootPlayer) GetValidMovesFor(cellName string, board Board) [] Move {
    return p.Pieces[cellName].GetValidMoves(board)
}


func (p* RootPlayer) UpdateCell(fromCellName, toCellName string, newCell Cell)  {
    piece := p.Pieces[fromCellName]
    delete(p.Pieces, fromCellName)
    piece.GetCommon().Cell = newCell
    p.Pieces[toCellName] = piece
}


func (p* RootPlayer) Remove(cellName string)  {
    delete(p.Pieces, cellName)
}

type Player interface {
    GetCommon() *RootPlayer
}

type WhitePlayer struct {
    Common RootPlayer
}

func (p WhitePlayer) GetCommon() *RootPlayer {
    return &p.Common
}

func MakeWhitePlayer(board Board) WhitePlayer {
    pieces := InitWhitePieces(board)
    return WhitePlayer{Common: RootPlayer{Pieces: pieces}}
}


type BlackPlayer struct {
    Common RootPlayer
}

func (p BlackPlayer) GetCommon() *RootPlayer {
    return &p.Common
}


func MakeBlackPlayer(board Board) BlackPlayer {
    pieces := InitBlackPieces(board)
    return BlackPlayer{Common: RootPlayer{Pieces: pieces}}
}

