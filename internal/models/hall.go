package models

import (
	"time"
	"gorm.io/gorm"
)

type Hall struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	TheaterID uint `json:"theater_id" gorm:"not null"`
	Capacity int `json:"capacity" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	//relationships
	Theater Theater `json:"theater" gorm:"foreignKey:TheaterID"`
	Seats []Seat `json:"seats" gorm:"foreignKey:HallID"`
	Shows []Show `json:"shows" gorm:"foreignKey:HallID"`
}