package viewModels

type ProductListVM struct {
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}
type ProductUpdateVM struct {
	Id           uint64
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}

type ProductCreateVM struct {
	Name         string
	Quantity     int
	BuyingPrice  float64
	SellingPrice float64
	Weight       float64
	Stock        int
}
