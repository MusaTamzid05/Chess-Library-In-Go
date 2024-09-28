package main

import (
    "chess_recording/lib"
    "fmt"
)


func test() {
    board := lib.MakeBoard()
    board.ShowInfo()

}

func main() {
    board := lib.MakeBoard()
    board.Render()

    blackPieces := lib.InitBlackPlayers(board)

    for cellName, piece := range blackPieces {
        moves := piece.GetValidMoves()
        fmt.Print(cellName,"\n")
        for _, move := range moves {
            fmt.Print(move.GetString(),"  ")
        }

        fmt.Println("\n")

    }

}
