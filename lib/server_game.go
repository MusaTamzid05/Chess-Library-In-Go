package lib

type ServerGame struct {
    players [] Player
    board Board
    CurrentPlayerIndex int
}


func MakeServerGame() ServerGame {
    board := MakeBoard()

    return ServerGame {
        board: board,
        CurrentPlayerIndex: 0,
        players: []Player {
            MakeWhitePlayer(board), 
            MakeBlackPlayer(board)},
        }
}

func (s ServerGame) GetBoardMap() [8][8] string {
    boardMap := [8][8]string {}

    for r := 0; r < 8; r += 1 {
        for c := 0; c < 8; c += 1 {
            boardMap[r][c] = s.board.Cells[r][c].Symbol
        }
    }

    return boardMap

}


func (s ServerGame) GetValidMoveNames(cellName string) []string {

    if s.players[s.CurrentPlayerIndex].GetCommon().PieceExists(cellName) == false {
        return []string {}
    }

    validMoves := s.players[s.CurrentPlayerIndex].GetCommon().GetValidMovesFor(cellName, s.board)
    validMovePosNames := []string{}

    for _, validMove := range validMoves {
        validMovePosNames = append(validMovePosNames, validMove.GetString())
    }

    return validMovePosNames

}

func (s *ServerGame) Update(fromCellName, toCellName string) error {
    currentPlayerIndex := s.CurrentPlayerIndex

    newCell, pieceRemoved, err := s.board.UpdateBoard(
        fromCellName,
        toCellName)


    if err != nil {
        return err
    }


    s.players[currentPlayerIndex].GetCommon().UpdateCell(
        fromCellName,
        toCellName,
        newCell)


    nextPlayerIndex := (currentPlayerIndex + 1) % 2

    if pieceRemoved {
        s.players[nextPlayerIndex].GetCommon().Remove(
            toCellName)
    }

    s.CurrentPlayerIndex = nextPlayerIndex

    return nil
}
