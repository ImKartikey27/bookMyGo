package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type hallRepository struct {
    db *gorm.DB
}

func NewHallRepository(db *gorm.DB) HallRepository {
    return &hallRepository{db: db}
}

func (r *hallRepository) Create(hall *models.Hall) error {
    return r.db.Create(hall).Error
}

func (r *hallRepository) GetByID(id uint) (*models.Hall, error) {
    var hall models.Hall
    err := r.db.Preload("Theater").Preload("Seats").First(&hall, id).Error
    return &hall, err
}

func (r *hallRepository) GetByTheaterID(theaterID uint) ([]models.Hall, error) {
    var halls []models.Hall
    err := r.db.Preload("Seats").Where("theater_id = ?", theaterID).Find(&halls).Error
    return halls, err
}

func (r *hallRepository) Update(hall *models.Hall) error {
    return r.db.Save(hall).Error
}

func (r *hallRepository) Delete(id uint) error {
    return r.db.Delete(&models.Hall{}, id).Error
}