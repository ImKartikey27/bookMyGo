package routes

import (
	"bookmygo/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	theaterController := controllers.NewTheaterController()
	movieController := controllers.NewMovieController()

	api := r.Group("/api/v1")
	{
		theaters := api.Group("/theaters")
		{
			theaters.POST("/", theaterController.CreateTheater)
			theaters.GET("/", theaterController.GetAllTheaters)
			theaters.GET("/:id", theaterController.GetTheaterByID)
			theaters.DELETE("/:id", theaterController.DeleteTheater)
		}

		movies := api.Group("/movies")
		{
			movies.POST("/", movieController.CreateMovie)
			movies.GET("/", movieController.GetAllMovies)
			movies.GET("/:id", movieController.GetMovieByID)
			movies.PUT("/:id", movieController.UpdateMovie)
			movies.DELETE("/:id", movieController.DeleteMovie)
		}
	}
}