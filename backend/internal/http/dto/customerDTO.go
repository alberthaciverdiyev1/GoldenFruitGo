package dto

import "time"

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Surname string `json:"surname" binding:"required"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CustomerResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
