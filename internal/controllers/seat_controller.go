package controllers

import (
    "net/http"
    "strconv"

    "bookmygo/internal/database"
    "bookmygo/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type SeatController struct {
    db *gorm.DB
}

func NewSeatController() *SeatController {
    return &SeatController{
        db: database.GetDB(),
    }
}

func (sc *SeatController) CreateSeatsForHall(c *gin.Context) {
    hallID, _ := strconv.Atoi(c.Param("hallId"))
    
    var requestData struct {
        Rows    int `json:"rows"`
        Columns int `json:"columns"`
    }
    
    if err := c.ShouldBindJSON(&requestData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if seats already exist for this hall
    var existingSeats []models.Seat
    sc.db.Where("hall_id = ?", hallID).Find(&existingSeats)
    if len(existingSeats) > 0 {
        c.JSON(http.StatusConflict, gin.H{"error": "Seats already exist for this hall"})
        return
    }

    var seats []models.Seat
    rowNames := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
    
    for i := 0; i < requestData.Rows && i < len(rowNames); i++ {
        for j := 1; j <= requestData.Columns; j++ {
            seat := models.Seat{
                HallID:     uint(hallID),
                SeatNumber: strconv.Itoa(j),
                Row:        rowNames[i],
            }
            seats = append(seats, seat)
        }
    }

    if err := sc.db.Create(&seats).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Seats created successfully",
        "total_seats": len(seats),
    })
}
func (sc *SeatController) GetSeatsByHall(c *gin.Context) {
    hallID, _ := strconv.Atoi(c.Param("hallId"))
    var seats []models.Seat

    if err := sc.db.Where("hall_id = ?", hallID).Order("row ASC, seat_number ASC").Find(&seats).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "hall_id":     hallID,
        "total_seats": len(seats),
        "seats":       seats,
    })
}