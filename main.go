package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"net/http"
)

func handleRoot(c *gin.Context) {

	info := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "MazerRackham",
		Color:      "#888888",
		Head:       "default",
		Tail:       "default",
	}

	c.JSON(http.StatusOK, info)

}

func handleStart(c *gin.Context) {
	var gamestate GameState

	err := c.ShouldBind(&gamestate)
	if err != nil {
		log.Print(err)
	}

}
func handleEnd(c *gin.Context) {
	var gamestate GameState

	err := c.ShouldBind(&gamestate)
	if err != nil {
		log.Print(err)
	}

}

func handleMove(c *gin.Context) {
	var gamestate GameState

	err := c.ShouldBind(&gamestate)
	if err != nil {
		log.Print(err)
		return
	}

	move := move(gamestate)

	c.JSON(http.StatusOK, move)

}

func main() {
	router := gin.Default()
	router.GET("/", handleRoot)
	router.POST("/start", handleStart)
	router.POST("/move", handleMove)
	router.POST("/end", handleEnd)

	router.Run("127.0.0.1:8080")
}
