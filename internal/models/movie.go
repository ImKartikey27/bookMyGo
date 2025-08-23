package models

import (
	"time"
	"gorm.io/gorm"
)

type Movie struct{
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Genre string `json:"genre" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Shows []Show `json:"shows" gorm:"foreignKey:MovieID"`
}