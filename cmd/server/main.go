package main

import (
	"net/http"
	"os"

	"bookmygo/internal/config"
	"bookmygo/internal/database"
	"bookmygo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode for production
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Load configuration
	cfg := config.LoadConfig()
	//connect to database
	database.ConnectDB(cfg)
	//run database migrations
	database.RunMigrations()
	//gin router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to BookMyGo API",
			"status":  "server running",
			"version": "v1.0",
		})
	})

	//health check route

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})
	//setup api routes
	routes.SetupRoutes(r)

	r.Run(":" + cfg.ServerPort)
}