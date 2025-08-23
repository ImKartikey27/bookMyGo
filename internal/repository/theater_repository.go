package repository

import (
	"bookmygo/internal/models"
)

type TheaterRepository interface {
	Create(theater *models.Theater) error
	GetByID(id uint) (*models.Theater, error)
	GetAll() ([]models.Theater, error)
	Update(theater *models.Theater) error
	Delete(id uint) error
}