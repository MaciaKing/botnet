package main

import (
	"botnet/database"
	"botnet/handler"
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

	// r.LoadHTMLGlob("./templates/*")
	r.LoadHTMLGlob("./templates/*.html")

	// Servir archivos estáticos desde `templates/assets`
	r.Static("/assets", "./templates/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
