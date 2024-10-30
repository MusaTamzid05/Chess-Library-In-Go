package main

import (
    "chess_recording/lib"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Server() {
    serverGame := lib.MakeServerGame()
    r := gin.Default()

    r.GET("/board", func(c *gin.Context) {
        boardMap := serverGame.GetBoardMap()
        c.JSON(http.StatusOK, boardMap)

    })

    r.Run(":8080")
}
