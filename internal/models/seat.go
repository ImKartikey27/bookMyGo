package models

import (
    "time"
    "gorm.io/gorm"
)

type Seat struct {
    ID uint `json:"id" gorm:"primaryKey"`
    HallID uint `json:"hall_id" gorm:"not null"`
    SeatNumber string `json:"seat_number" gorm:"not null"`
    Row string `json:"row" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
    
    // Relationships
    Hall Hall `json:"hall" gorm:"foreignKey:HallID"`
    Bookings []Booking `json:"bookings" gorm:"foreignKey:SeatID"`
}