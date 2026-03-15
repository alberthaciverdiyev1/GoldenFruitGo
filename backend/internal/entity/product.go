package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}
