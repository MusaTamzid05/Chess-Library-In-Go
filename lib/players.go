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
        pieces[string(CHARS[i]) + "7"] = MakePawn(board.Cells[1][7], false)
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
