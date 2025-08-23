package models

import (
    "time"
    "gorm.io/gorm"
)

type Show struct {
    ID uint `json:"id" gorm:"primaryKey"`
    MovieID uint `json:"movie_id" gorm:"not null"`
    HallID uint `json:"hall_id" gorm:"not null"`
    ShowTime time.Time `json:"show_time" gorm:"not null"`
    Price float64 `json:"price" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
    
    // Relationships
    Movie Movie `json:"movie" gorm:"foreignKey:MovieID"`
    Hall Hall `json:"hall" gorm:"foreignKey:HallID"`
    Bookings []Booking `json:"bookings" gorm:"foreignKey:ShowID"`
}