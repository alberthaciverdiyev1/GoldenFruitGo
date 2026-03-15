package dto

type ProductResponse struct {
	Id           uint64
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}
type UpdateProductRequest struct {
	Id           uint64
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}

type CreateProductRequest struct {
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}
