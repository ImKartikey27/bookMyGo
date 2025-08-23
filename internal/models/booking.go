package models

import (
    "time"
    "gorm.io/gorm"
)

type Booking struct {
    ID uint `json:"id" gorm:"primaryKey"`
    ShowID uint `json:"show_id" gorm:"not null"`
    SeatID uint `json:"seat_id" gorm:"not null"`
    CustomerName string `json:"customer_name" gorm:"not null"`
    CustomerEmail string `json:"customer_email" gorm:"not null"`
    CustomerPhone string `json:"customer_phone" gorm:"not null"`
    IsBooked bool `json:"is_booked" gorm:"default:false"`
    BookingTime time.Time `json:"booking_time"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
    
    Show Show `json:"show" gorm:"foreignKey:ShowID"`
    Seat Seat `json:"seat" gorm:"foreignKey:SeatID"`
}