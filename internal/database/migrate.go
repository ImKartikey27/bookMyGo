package database

import (
	"log"
	"bookmygo/internal/models"
)

func RunMigrations(){
	db := GetDB()
	log.Println("Running Database Migrations")

	err := db.AutoMigrate(
		&models.Theater{},
		&models.Hall{},
		&models.Movie{},
		&models.Show{},
		&models.Seat{},
		&models.Booking{},
	)

	if err != nil{
		log.Fatal("Failed to run database migrations")
	}
	log.Println("Database Migrations completed successfully")
}