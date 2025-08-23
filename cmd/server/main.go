package main

import (
	"net/http"

	"bookmygo/internal/config"
	"github.com/gin-gonic/gin"
)

func main(){
	//Load configuration
	cfg := config.LoadConfig()
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

	r.Run(":"+ cfg.SERVER_PORT)
}