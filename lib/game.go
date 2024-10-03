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


        validMove := false
        fromCellName := ""
        toCellName := ""

        for validMove == false {
            fmt.Printf(">>> ")
            fmt.Scan(&userInput)

            if userInput == "exit" {
                running = false
                break
            }

            fromCellName = userInput[0:2]
            toCellName = userInput[2:4]

            fmt.Println(fromCellName, toCellName)


            // convert the move into board cell map

            newMove, err := MakeMove(toCellName)
            oldMove , _ := MakeMove(fromCellName)

            if err != nil {
                fmt.Println(err)
                continue
            }


            fmt.Println(oldMove.Row, oldMove.Col)
            fmt.Println(newMove.Row, newMove.Col)
        }

        if running == false {
            continue
        }




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

