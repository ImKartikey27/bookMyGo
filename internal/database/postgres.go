package database

import (
	"fmt"
	"log"

	"bookmygo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB*gorm.DB

func ConnectDB(cfg *config.Config) {
	var dsn string
	
	// Use DATABASE_URL if available (Render provides this)
	if cfg.DatabaseURL != "" {
		dsn = cfg.DatabaseURL
	} else {
		// Fallback to individual components
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully!")
}

func GetDB() *gorm.DB {
	return DB
}