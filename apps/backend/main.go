package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/santiagofelizzola/coaching-app/database"
	"github.com/santiagofelizzola/coaching-app/routes"
	"github.com/santiagofelizzola/coaching-app/config"
)

func main() {
	// Load up env & connect to database
	config.LoadEnv()
	database.Connect()

	// Init Gin & Routes
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	routes.RegisterRoutes(r)

	// Start server
	port := config.GetPort()
	log.Printf("Listening on port %s...", port)
	r.Run(":" + port)
}
