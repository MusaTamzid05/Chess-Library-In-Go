package lib

import (
    "fmt"
)

type Game struct {
    players [] Player
    board Board
}

func (g* Game) Run() {
    running := true
    var userInput string

    for running {
        g.board.Render()


        validMoveFound := false
        fromCellName := ""
        toCellName := ""
        playerColors := [] string {"white", "black"}
        currentPlayerIndex := 0

        for validMoveFound == false {
            fmt.Printf("%s >>> ", playerColors[currentPlayerIndex])
            fmt.Scan(&userInput)

            if userInput == "exit" {
                running = false
                break
            }

            fromCellName = userInput[0:2]
            toCellName = userInput[2:4]

            fmt.Println(fromCellName, toCellName)


            if g.players[currentPlayerIndex].GetCommon().PieceExists(fromCellName) == false {
                fmt.Println("Invalid Piece for the player ", fromCellName)
                continue
            }

            toMove, err := MakeMove(toCellName)

            if err != nil {
                fmt.Println(err)
                continue
            }

            validMoves := g.players[currentPlayerIndex].GetCommon().GetValidMovesFor(fromCellName)

            for _, validMove := range validMoves {
                if validMove.Equal(toMove) {
                    validMoveFound = true
                    break
                }

            }

            if validMoveFound == false {
                fmt.Println(toCellName, " is not a valid move")
            }

        }

        fmt.Println("valid move found")

        if running == false {
            continue
        }

        nextPlayerIndex := (currentPlayerIndex + 1) % 2
        currentPlayerIndex = nextPlayerIndex




    }

}

func MakeGame() Game {
    board := MakeBoard()

    return Game{
        board: board,
        players: []Player {
        MakeWhitePlayer(board), 
        MakeBlackPlayer(board),
    }}
}

