package controllers

import (
	"net/http"
	"strconv"

	"bookmygo/internal/database"
	"bookmygo/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShowController struct {
	db *gorm.DB
}

func NewShowController() *ShowController{
	return &ShowController{
		db: database.GetDB(),
	}
}

func (sc *ShowController) CreateShow(c *gin.Context){
	var show models.Show
	if err:= c.ShouldBindJSON(&show); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.db.Create(&show).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, show)
}

func (sc *ShowController) GetAllShows(c *gin.Context){
	var shows []models.Show
	if err := sc.db.Preload("Movie").Preload("Hall.Theater").Find(&shows).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shows)
}

func (sc *ShowController) GetShowByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var show models.Show

	if err := sc.db.Preload("Movie").Preload("Hall.Theater").First(&show, id).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Show not found"})
		return
	}

	c.JSON(http.StatusOK, show)
}

func (sc *ShowController) GetShowsByMovie(c *gin.Context){
	movieID, _ := strconv.Atoi(c.Param("movieID"))
	var shows []models.Show

	if err := sc.db.Preload("Movie").Preload("Hall.Theater").Where("movie_id = ?", movieID).Find(&shows).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shows not found"})
		return
	}

	c.JSON(http.StatusOK, shows)
}

func (sc *ShowController) GetShowsByTheater(c *gin.Context){
	theaterID, _ := strconv.Atoi(c.Param("theaterID"))
	var shows []models.Show

	if err := sc.db.Preload("Movie").Preload("Hall.Theater").Joins("JOIN halls ON shows.hall_id = halls.id").Where("halls.theater_id = ?", theaterID).Find(&shows).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shows not found"})
		return
	}

	c.JSON(http.StatusOK, shows)
}

func (sc *ShowController) DeleteShow(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	if err := sc.db.Delete(&models.Show{}, id).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Show deleted successfully"})
}