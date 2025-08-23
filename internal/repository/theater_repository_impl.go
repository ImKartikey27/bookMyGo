package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type theaterRepository struct {
    db *gorm.DB
}

func NewTheaterRepository(db *gorm.DB) TheaterRepository {
    return &theaterRepository{db: db}
}

func (r *theaterRepository) Create(theater *models.Theater) error {
    return r.db.Create(theater).Error
}

func (r *theaterRepository) GetByID(id uint) (*models.Theater, error) {
    var theater models.Theater
    err := r.db.Preload("Halls").First(&theater, id).Error
    return &theater, err
}

func (r *theaterRepository) GetAll() ([]models.Theater, error) {
    var theaters []models.Theater
    err := r.db.Preload("Halls").Find(&theaters).Error
    return theaters, err
}

func (r *theaterRepository) Update(theater *models.Theater) error {
    return r.db.Save(theater).Error
}

func (r *theaterRepository) Delete(id uint) error {
    return r.db.Delete(&models.Theater{}, id).Error
}