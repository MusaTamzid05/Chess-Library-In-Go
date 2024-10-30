package main

import (
    "chess_recording/lib"
    //"fmt"
)


func test() {
    board := lib.MakeBoard()
    board.ShowInfo()

}

func main() {
    /*
    board := lib.MakeBoard()
    board.Render()

    whitePieces := lib.InitWhitePieces(board)
    bishop := whitePieces["f1"]

    for cellName, move := range bishop.GetValidMoves() {
        fmt.Println(cellName, move.GetString())
    }
    */

    //game := lib.MakeGame()
    //game.Run()

    Server()


}
