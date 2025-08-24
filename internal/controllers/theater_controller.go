package controllers

import (
	"net/http"
	"strconv"

	"bookmygo/internal/database"
	"bookmygo/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TheaterController struct {
	db *gorm.DB
}

func NewTheaterController() *TheaterController {
	return &TheaterController{
		db: database.GetDB(),
	}
}

func (tc *TheaterController) CreateTheater(c *gin.Context){
	var theater models.Theater
	if err := c.ShouldBindJSON(&theater); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.db.Create(&theater).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, theater)
}

func (tc *TheaterController) GetAllTheaters(c *gin.Context){
	var theaters []models.Theater
	if err := tc.db.Find(&theaters).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, theaters)
}

func (tc *TheaterController) GetTheaterByID(c* gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var theater models.Theater

	if err := tc.db.First(&theater, id).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Theater not found"})
		return
	}

	c.JSON(http.StatusOK, theater)
}