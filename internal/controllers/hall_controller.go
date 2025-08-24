package controllers

import (
	"net/http"
	"strconv"

	"bookmygo/internal/database"
	"bookmygo/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HallController struct {
	db *gorm.DB
}

func NewHallController() *HallController {
	return &HallController{
		db: database.GetDB(),
	}
}

func (hc *HallController) CreateHall(c *gin.Context){
	var hall models.Hall

	if err := c.ShouldBindJSON(&hall); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hc.db.Create(&hall).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hall)
}

func (hc *HallController) GetAllHalls(c *gin.Context){
	var halls []models.Hall

	if err := hc.db.Preload("Theater").Find(&halls).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, halls)
}

func (hc *HallController) GetHallByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var hall models.Hall

	if err := hc.db.Preload("Theater").Preload("Seats").Find(&hall, id).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hall)
}

func (hc *HallController) DeleteHall(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	if err := hc.db.Delete(&models.Hall{}, id).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hall deleted successfully"})
}