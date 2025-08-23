package main

import (
	"net/http"

	"bookmygo/internal/database"
	"bookmygo/internal/config"
	"github.com/gin-gonic/gin"
)

func main(){
	//Load configuration
	cfg := config.LoadConfig()
	//connect to database
	database.ConnectDB(cfg)
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

	r.GET("/config", func(c*gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"DBHost": cfg.DBHost,
			"DBPort": cfg.DBPort,
			"DBUser": cfg.DBUser,
			"DBName": cfg.DBName,
			"SERVER_PORT": cfg.SERVER_PORT,
		})
	})

	r.GET("/db-test", func(c*gin.Context){
		db := database.GetDB()
		_  , err := db.DB()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Database connection failed",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Database connection successful",
		})
	})

	r.Run(":"+ cfg.SERVER_PORT)
}