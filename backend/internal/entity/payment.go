package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentCategory string

const (
	CategoryIncome  PaymentCategory = "INCOME"
	CategoryExpense PaymentCategory = "EXPENSE"
)

type Payment struct {
	gorm.Model
	CustomerID  uint            `gorm:"not null;index"`
	Customer    Customer        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderID     uint64          `gorm:"index"`
	Amount      int64           `gorm:"not null"`
	Category    PaymentCategory `gorm:"size:10;not null"`
	ProcessedAt *time.Time      `gorm:"index"`
}
