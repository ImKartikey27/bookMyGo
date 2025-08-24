package routes

import (
	"bookmygo/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	theaterController := controllers.NewTheaterController()
	movieController := controllers.NewMovieController()
	bookingController := controllers.NewBookingController()
	showController := controllers.NewShowController()
	seatController := controllers.NewSeatController()
	hallController := controllers.NewHallController()

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

		booking := api.Group("/bookings")
		{
			booking.POST("/", bookingController.CreateBooking)
			booking.GET("/:id", bookingController.GetBookingByID)
			booking.GET("/show/:showID", bookingController.GetBookingsByShow)
			booking.PUT("/:id", bookingController.CancelBooking)
			booking.GET("/check-availability", bookingController.CheckSeatAvailability)
			booking.GET("/available-seats/:showID", bookingController.GetAvailableSeats)

		}
		shows := api.Group("/shows")
		{
			shows.POST("/", showController.CreateShow)
			shows.GET("/", showController.GetAllShows)
			shows.GET("/:id", showController.GetShowByID)
			shows.GET("/movie/:movieID", showController.GetShowsByMovie)
			shows.GET("/theater/:theaterID",showController.GetShowsByTheater)
			shows.DELETE("/:id", showController.DeleteShow)
		}
		seats := api.Group("/seats")
		{
			seats.POST("/hall/:hallId", seatController.CreateSeatsForHall)
			seats.GET("/hall/:hallId", seatController.GetSeatsByHall)
		}
		halls := api.Group("/halls")
		{
			halls.POST("/", hallController.CreateHall)
			halls.GET("/", hallController.GetAllHalls)
			halls.GET("/:id", hallController.GetHallByID)
			halls.DELETE("/:id", hallController.DeleteHall)
		}


		
	}
}