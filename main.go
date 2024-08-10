package main

import (
	"botnet/database"
	"botnet/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bots", handler.GetAllBots)
	r.Run() // listen and serve on 0.0.0.0:8080
}
