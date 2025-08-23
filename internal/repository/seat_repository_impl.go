package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type seatRepository struct {
    db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) SeatRepository {
    return &seatRepository{db: db}
}

func (r *seatRepository) Create(seat *models.Seat) error {
    return r.db.Create(seat).Error
}

func (r *seatRepository) GetByID(id uint) (*models.Seat, error) {
    var seat models.Seat
    err := r.db.First(&seat, id).Error
    return &seat, err
}

func (r *seatRepository) GetByHallID(hallID uint) ([]models.Seat, error) {
    var seats []models.Seat
    err := r.db.Where("hall_id = ?", hallID).Find(&seats).Error
    return seats, err
}

func (r *seatRepository) Update(seat *models.Seat) error {
    return r.db.Save(seat).Error
}

func (r *seatRepository) Delete(id uint) error {
    return r.db.Delete(&models.Seat{}, id).Error
}

func (r *seatRepository) CreateMultiple(seats []models.Seat) error {
    return r.db.Create(&seats).Error
}