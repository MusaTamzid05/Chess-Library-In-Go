package lib

func InitBlackPlayers(board Board) map[string]Piece {
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
