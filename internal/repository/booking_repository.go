package repository

import "bookmygo/internal/models"

type BookingRepository interface {
    Create(booking *models.Booking) error
    GetByID(id uint) (*models.Booking, error)
    GetByShowID(showID uint) ([]models.Booking, error)
    GetBySeatAndShow(seatID, showID uint) (*models.Booking, error)
    Update(booking *models.Booking) error
    Delete(id uint) error
    IsSeatBooked(seatID, showID uint) (bool, error)
}