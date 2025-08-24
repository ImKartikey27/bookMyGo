package controllers

import (
	"net/http"
	"strconv"

	"bookmygo/internal/database"
	"bookmygo/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController struct {
	db *gorm.DB
}

func NewMovieController() *MovieController {
	return &MovieController{
		db: database.GetDB(),
	}
}

func (mc *MovieController) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	if err := mc.db.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
	 

}

func (mc *MovieController) GetAllMovies(c *gin.Context){
	var movies []models.Movie
	if err := mc.db.Find(&movies).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (mc *MovieController) GetMovieByID(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var movie models.Movie

	if err := mc.db.First(&movie, id).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

//update will work with all field otherwise makes it empty
func (mc* MovieController) UpdateMovie(c* gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var movie models.Movie

	if err := mc.db.First(&movie, id).Error; err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := mc.db.Save(&movie).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (mc* MovieController) DeleteMovie(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	if err := mc.db.Delete(&models.Movie{}, id).Error; err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie deleted successfully"})
}