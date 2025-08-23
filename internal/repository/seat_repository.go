package repository

import "bookmygo/internal/models"

type SeatRepository interface {
    Create(seat *models.Seat) error
    GetByID(id uint) (*models.Seat, error)
    GetByHallID(hallID uint) ([]models.Seat, error)
    Update(seat *models.Seat) error
    Delete(id uint) error
    CreateMultiple(seats []models.Seat) error
}