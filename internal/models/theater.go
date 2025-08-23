package models

import (
	"time"
	"gorm.io/gorm"
)

type Theater struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
	City string `json:"city" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Halls []Hall `json:"halls" gorm:"foreignKey:TheaterID"` 
 }