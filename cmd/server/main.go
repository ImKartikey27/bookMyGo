package main

import (
	"net/http"

	"bookmygo/internal/database"
	"bookmygo/internal/config"
	"bookmygo/internal/models"
	"bookmygo/internal/repository"
	"github.com/gin-gonic/gin"
)

func main(){
	//Load configuration
	cfg := config.LoadConfig()
	//connect to database
	database.ConnectDB(cfg)
	//run database migrations
	database.RunMigrations()
	//initialize repositories
	repos := repository.NewRepositories(database.GetDB())
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

	r.GET("/test-repos" ,func(c*gin.Context){
		//test theater repository
		theater := &models.Theater{
			Name:"Test Cinema2",
			Address:"123 Test Street",
			City:"Test City",
		}
		if err := repos.Theater.Create(theater); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed to create theater",
				"error": err,
			})
			return

		}
		//test movies
		movie := &models.Movie{
			Title:"Test Movie",
			Description:"Test Description",
			Genre:"Test Genre",
		}
		if err := repos.Movie.Create(movie); err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed to create movie",
				"error": err,
			})
			return
		}
		//get all data
		theaters, err := repos.Theater.GetAll()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed to get theaters",
				"error": err,
			})
			return
		}
		//get all movies
		movies, err := repos.Movie.GetAll()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed to get movies",
				"error": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Test successful",
			"theaters": theaters,
			"movies": movies,
		})
		
	})

	r.Run(":"+ cfg.SERVER_PORT)
}