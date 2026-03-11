package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string    `gorm:"not null;type:varchar(255)"`
	Surname  string    `gorm:"default:null;type:varchar(255)"`
	Phone    string    `gorm:"default:null;type:varchar(255)"`
	Address  string    `gorm:"default:null;type:varchar(255)"`
	Payments []Payment `gorm:"foreignKey:CustomerID"`
}
