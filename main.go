package main

import (
	"botnet/database"
	"botnet/handler"
	"botnet/models"
	"net/http"

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

	r.LoadHTMLGlob("./templates/*.html")
	r.Static("/assets", "./templates/assets")

	r.GET("/", func(c *gin.Context) {
		// get all bots
		var bots []models.Bot
		if err := database.DB.Find(&bots).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"bots": bots,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
