package controllers

import (
	"net/http"
	"strconv"
	"time"

	"bookmygo/internal/database"
	"bookmygo/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingController struct {
	db *gorm.DB
}


func NewBookingController() *BookingController{
	return &BookingController{
		db: database.GetDB(),
	}
}
func (bc *BookingController) GetAvailableSeats(c *gin.Context){
	showID, _ := strconv.Atoi(c.Param("showId"))
	if showID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "show_id is required"})
		return
	}
	//get all booked seat IDs for the show
	var bookedseatIDs []uint 
	if err := bc.db.Model(&models.Booking{}).Where("show_id = ? AND is_booked = ?", showID, true).Pluck("seat_id", &bookedseatIDs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//get show details to find the hall
	var show models.Show
	if err := bc.db.Preload("Hall.Theater").First(&show, showID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Show not found"})
		return
	}
	//get all seats for this hall 
	var allSeats []models.Seat
	query := bc.db.Preload("Hall.Theater").Where("hall_id = ?", show.HallID)

	//exclude booked seats
	if len(bookedseatIDs) > 0 {
		query = query.Where("id NOT IN ?", bookedseatIDs)
	}

	if err := query.Find(&allSeats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"show_id": showID,
		"hall_name": show.Hall.Name,
		"theater_name": show.Hall.Theater.Name,
		"available_seats": allSeats,
		"total_available": len(allSeats),
	})

}

func (bc *BookingController) CreateBooking(c *gin.Context){
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//start database transaction 
	tx := bc.db.Begin()
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//check if seat is already booked

	var existingBooking models.Booking
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("seat_id = ? AND show_id = ? AND is_booked = ?", booking.SeatID, booking.ShowID, true).First(&existingBooking).Error; err == nil{
		tx.Rollback()
		c.JSON(http.StatusConflict, gin.H{"error": "Seat is already booked"})
		return
	}

	booking.IsBooked = true
	booking.BookingTime = time.Now()

	//create booking within the transaction
	if err := tx.Create(&booking).Error; err!= nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//commit transaction
	if err := tx.Commit().Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fetch preloads
	var createdBooking models.Booking
	if err := bc.db.Preload("Show.Movie").Preload("Show.Hall.Theater").Preload("Seat.Hall").First(&createdBooking, booking.ID).Error; err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
    "message": "Booking created successfully",
    "booking_details": gin.H{
        "booking_id":    createdBooking.ID,
        "customer_name": createdBooking.CustomerName,
        "movie_title":   createdBooking.Show.Movie.Title,
        "theater_name":  createdBooking.Show.Hall.Theater.Name,
        "hall_name":     createdBooking.Show.Hall.Name,
        "seat":          createdBooking.Seat.Row + createdBooking.Seat.SeatNumber,
        "show_time":     createdBooking.Show.ShowTime,
        "price":         createdBooking.Show.Price,
        "booking_time":  createdBooking.BookingTime,
    },
})
}
func (bc *BookingController) GetBookingByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var booking models.Booking

	if err := bc.db.Preload("Show.Movie").Preload("Show.Hall.Theater").Preload("Seat").First(&booking, id).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
    "message": "Booking created successfully",
    "booking_details": gin.H{
        "booking_id":    booking.ID,
		"booking_status": booking.IsBooked,
        "customer_name": booking.CustomerName,
        "movie_title":   booking.Show.Movie.Title,
        "theater_name":  booking.Show.Hall.Theater.Name,
        "hall_name":     booking.Show.Hall.Name,
        "seat":          booking.Seat.Row + booking.Seat.SeatNumber,
        "show_time":     booking.Show.ShowTime,
        "price":         booking.Show.Price,
        "booking_time":  booking.BookingTime,
    },
})
}
func (bc *BookingController) CancelBooking(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var booking models.Booking

	if err := bc.db.First(&booking, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}
	booking.IsBooked = false
	if err := bc.db.Save(&booking).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking canceled successfully"})
}