package repository

import "bookmygo/internal/models"

type ShowRepository interface {
    Create(show *models.Show) error
    GetByID(id uint) (*models.Show, error)
    GetAll() ([]models.Show, error)
    GetByMovieID(movieID uint) ([]models.Show, error)
    GetByTheaterID(theaterID uint) ([]models.Show, error)
    Update(show *models.Show) error
    Delete(id uint) error
}