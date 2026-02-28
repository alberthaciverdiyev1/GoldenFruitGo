package viewModels

import "time"

type CreateCustomerVM struct {
	Name    string `json:"name" binding:"required,min=2"`
	Surname string `json:"surname" binding:"omitempty,min=2"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateCustomerVM struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CustomerResponseVM struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Debt      float64   `json:"debt"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerPurchaseVM struct {
	ProductName string    `json:"product_name"`
	Quantity    float64   `json:"quantity"`
	Price       float64   `json:"price"`
	TotalPrice  float64   `json:"total_price"`
	Date        time.Time `json:"date"`
}

type CustomerDebtLogVM struct {
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
}

type CustomerDetailsVM struct {
	ID        uint64               `json:"id"`
	Name      string               `json:"name"`
	Surname   string               `json:"surname"`
	Phone     string               `json:"phone"`
	Address   string               `json:"address"`
	Debt      float64              `json:"debt"`
	Purchases []CustomerPurchaseVM `json:"purchases"`
	DebtLogs  []CustomerDebtLogVM  `json:"debt_logs"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
