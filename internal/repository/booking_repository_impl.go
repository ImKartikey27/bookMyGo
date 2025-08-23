package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type bookingRepository struct {
    db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
    return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(booking *models.Booking) error {
    return r.db.Create(booking).Error
}

func (r *bookingRepository) GetByID(id uint) (*models.Booking, error) {
    var booking models.Booking
    err := r.db.Preload("Show.Movie").Preload("Show.Hall.Theater").Preload("Seat").First(&booking, id).Error
    return &booking, err
}

func (r *bookingRepository) GetByShowID(showID uint) ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.db.Preload("Seat").Where("show_id = ? AND is_booked = ?", showID, true).Find(&bookings).Error
    return bookings, err
}

func (r *bookingRepository) GetBySeatAndShow(seatID, showID uint) (*models.Booking, error) {
    var booking models.Booking
    err := r.db.Where("seat_id = ? AND show_id = ? AND is_booked = ?", seatID, showID, true).First(&booking).Error
    return &booking, err
}

func (r *bookingRepository) Update(booking *models.Booking) error {
    return r.db.Save(booking).Error
}

func (r *bookingRepository) Delete(id uint) error {
    return r.db.Delete(&models.Booking{}, id).Error
}

func (r *bookingRepository) IsSeatBooked(seatID, showID uint) (bool, error) {
    var count int64
    err := r.db.Model(&models.Booking{}).Where("seat_id = ? AND show_id = ? AND is_booked = ?", seatID, showID, true).Count(&count).Error
    return count > 0, err
}