package database

import (
	"log"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("goldenfruit.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot Connect To Database: %v", err)
	}

	err = db.AutoMigrate(&entity.Customer{}, &entity.User{})
	if err != nil {
		log.Fatalf("Migration Error: %v", err)
	}

	SeedDefaultUser(db)
	return db
}
