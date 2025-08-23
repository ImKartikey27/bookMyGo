package repository

import "bookmygo/internal/models"

type HallRepository interface {
    Create(hall *models.Hall) error
    GetByID(id uint) (*models.Hall, error)
    GetByTheaterID(theaterID uint) ([]models.Hall, error)
    Update(hall *models.Hall) error
    Delete(id uint) error
}