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


    r.POST("/update", func(c *gin.Context) {
        var parameter struct {
            From string `json:"from"`
            To string `json:"to"`
        }

        if err := c.ShouldBindBodyWithJSON(&parameter); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
            return
        }

        validMovePosNames := serverGame.GetValidMoveNames(parameter.From)

        if len(validMovePosNames) == 0 {
            c.AsciiJSON(http.StatusOK, map[string]string{"status" : "failure", "message" : "Not Valid"})
            return

        }

        err := serverGame.Update(parameter.From, parameter.To)

        if err != nil {
            c.AsciiJSON(http.StatusOK, map[string]string{"status" : "failure", "message" : err.Error()})
            return

        }

        c.AsciiJSON(http.StatusOK, map[string]string{"status" : "success", "message" : ""})

    })


    r.Run(":8080")
}
