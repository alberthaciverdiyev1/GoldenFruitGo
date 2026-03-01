package viewModels

import "time"

type PurchaseVM struct {
	Id          uint64
	Customer    CustomerResponseVM
	CrateDate   time.Time
	Quantity    int64
	Weight      float64
	TotalWeight float64
	NetWeight   float64
	TotalPrice  float64
}

type PurchaseRequestVM struct {
	Id          uint64
	CustomerID  uint64
	Quantity    int64
	Weight      float64
	TotalWeight float64
	NetWeight   float64
	TotalPrice  float64
}
