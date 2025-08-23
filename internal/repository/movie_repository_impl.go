package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type movieRepository struct {
    db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
    return &movieRepository{db: db}
}

func (r *movieRepository) Create(movie *models.Movie) error {
    return r.db.Create(movie).Error
}

func (r *movieRepository) GetByID(id uint) (*models.Movie, error) {
    var movie models.Movie
    err := r.db.First(&movie, id).Error
    return &movie, err
}

func (r *movieRepository) GetAll() ([]models.Movie, error) {
    var movies []models.Movie
    err := r.db.Find(&movies).Error
    return movies, err
}

func (r *movieRepository) Update(movie *models.Movie) error {
    return r.db.Save(movie).Error
}

func (r *movieRepository) Delete(id uint) error {
    return r.db.Delete(&models.Movie{}, id).Error
}