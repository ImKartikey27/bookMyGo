package main

import (
	"net/http"

	"bookmygo/internal/database"
	"bookmygo/internal/config"
	"bookmygo/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	//Load configuration
	cfg := config.LoadConfig()
	//connect to database
	database.ConnectDB(cfg)
	//run database migrations
	database.RunMigrations()
	//gin router
	r := gin.Default()

	r.GET("/", func(c*gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"status": "server running",
		})
	})

	//health check route

	r.GET("/health", func(c*gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "Healthy",
		})
	})
	//setup api routes
	routes.SetupRoutes(r)


	r.Run(":"+ cfg.SERVER_PORT)
}