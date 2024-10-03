package lib

import (
    "fmt"
    "os"
    "os/exec"
)

type Game struct {
    players [] Player
    board Board
}

func (g* Game) Run() {
    running := true


    var userInput string
    currentPlayerIndex := 0
    fromCellName := ""
    toCellName := ""
    playerColors := [] string {"white", "black"}


    for running {
        g.board.Render()


        validMoveFound := false

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
                fmt.Println("invalid move")
            }

        }


        if running == false {
            continue
        }


        newCell, pieceRemoved, err := g.board.UpdateBoard(
            fromCellName,
            toCellName)


        if err != nil {
            fmt.Println(err)
            continue
        }


        g.players[currentPlayerIndex].GetCommon().UpdateCell(
            fromCellName,
            toCellName,
            newCell)


        nextPlayerIndex := (currentPlayerIndex + 1) % 2

        if pieceRemoved {
            g.players[nextPlayerIndex].GetCommon().Remove(
                toCellName)
        }

        currentPlayerIndex = nextPlayerIndex
        g.ClearScreen()



    }

}

func (g Game) ClearScreen() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()

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

