package main

import (
    "chess_recording/lib"
)


func test() {
    board := lib.MakeBoard()
    board.ShowInfo()




}

func main() {
    board := lib.MakeBoard()
    board.Render()
}
