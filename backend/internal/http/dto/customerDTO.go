package dto

import "time"

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Surname string `json:"surname" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email" binding:"omitempty,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Image   string `json:"image"`
}

type CustomerResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
