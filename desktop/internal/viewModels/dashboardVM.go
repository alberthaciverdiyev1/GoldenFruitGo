package viewModels

type DashboardVM struct {
	TotalSalesToday     float64
	TotalPurchasesToday float64
	TotalStockWeight    float64
	ActiveCustomerCount int
	RecentTransactions  []TransactionVM
	LowStockProducts    []ProductListVM
}

type TransactionVM struct {
	Type   string
	Party  string
	Amount float64
	Time   string
}
