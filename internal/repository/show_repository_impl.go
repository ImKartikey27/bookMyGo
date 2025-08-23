package repository

import (
    "bookmygo/internal/models"
    "gorm.io/gorm"
)

type showRepository struct {
    db *gorm.DB
}

func NewShowRepository(db *gorm.DB) ShowRepository {
    return &showRepository{db: db}
}

func (r *showRepository) Create(show *models.Show) error {
    return r.db.Create(show).Error
}

func (r *showRepository) GetByID(id uint) (*models.Show, error) {
    var show models.Show
    err := r.db.Preload("Movie").Preload("Hall.Theater").First(&show, id).Error
    return &show, err
}

func (r *showRepository) GetAll() ([]models.Show, error) {
    var shows []models.Show
    err := r.db.Preload("Movie").Preload("Hall.Theater").Find(&shows).Error
    return shows, err
}

func (r *showRepository) GetByMovieID(movieID uint) ([]models.Show, error) {
    var shows []models.Show
    err := r.db.Preload("Movie").Preload("Hall.Theater").Where("movie_id = ?", movieID).Find(&shows).Error
    return shows, err
}

func (r *showRepository) GetByTheaterID(theaterID uint) ([]models.Show, error) {
    var shows []models.Show
    err := r.db.Preload("Movie").Preload("Hall.Theater").
        Joins("JOIN halls ON shows.hall_id = halls.id").
        Where("halls.theater_id = ?", theaterID).Find(&shows).Error
    return shows, err
}

func (r *showRepository) Update(show *models.Show) error {
    return r.db.Save(show).Error
}

func (r *showRepository) Delete(id uint) error {
    return r.db.Delete(&models.Show{}, id).Error
}