package repository

import "bookmygo/internal/models"

type MovieRepository interface {
    Create(movie *models.Movie) error
    GetByID(id uint) (*models.Movie, error)
    GetAll() ([]models.Movie, error)
    Update(movie *models.Movie) error
    Delete(id uint) error
}