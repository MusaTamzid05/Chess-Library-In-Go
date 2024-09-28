package main

import (
    "chess_recording/lib"
)

func main() {
    board := lib.MakeBoard()
    board.Render()
}
