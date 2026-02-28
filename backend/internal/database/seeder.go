package database

import (
	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedDefaultUser(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("polad123"), bcrypt.DefaultCost)

	defaultUser := entity.User{
		UserName: "Polad",
		Password: string(hashedPassword),
	}

	db.Where(entity.User{UserName: "Polad"}).FirstOrCreate(&defaultUser)
}
