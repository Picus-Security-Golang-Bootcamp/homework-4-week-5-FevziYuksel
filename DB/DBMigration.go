package DBPackage

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func InitialMigration() *gorm.DB {
	// DB connection and create repositories
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")
	return db
}
