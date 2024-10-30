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


    r.POST("/valid_moves", func(c *gin.Context) {
        var parameter struct {
            Cell string `json:"cell"`
        }

        if err := c.ShouldBindBodyWithJSON(&parameter); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
            return
        }

        validMovePosNames := serverGame.GetValidMoveNames(parameter.Cell)
        c.JSON(http.StatusOK, validMovePosNames)
    })

    r.Run(":8080")
}
