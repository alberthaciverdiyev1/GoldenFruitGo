package main

import (
	"bytes"
	"context"
	"desktop/internal/api"
	"desktop/internal/handlers"
	"desktop/internal/viewModels"
	"desktop/internal/views/customer"
	"desktop/internal/views/dashboard"
	"desktop/internal/views/product"
	purchases "desktop/internal/views/purchase"
	sales "desktop/internal/views/sale"
	"fmt"
	"time"
)

type App struct {
	ctx  context.Context
	Auth *handlers.AuthHandler
}

func NewApp() *App {
	apiClient := api.NewClient("http://localhost:8080/api/v1")

	return &App{
		Auth: &handlers.AuthHandler{API: apiClient},
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetLoginPageHTML(errMsg string) string {
	return a.Auth.GetLoginPageHTML(errMsg)
}

func (a *App) DoLogin(user, pass string) viewModels.LoginResult {
	res, err := a.Auth.Login(user, pass)

	if err != nil {
		return viewModels.LoginResult{Success: false, Message: err.Error()}
	}

	return viewModels.LoginResult{
		Success: true,
		Message: "Giriş uğurludur",
		User:    res.UserName,
		Token:   res.Token,
	}
}

func (a *App) GetProductListHTML() string {
	// 1. API-dan məhsulları gətiririk (Sizin mövcud API Client ilə)
	// products, err := a.Auth.API.GetProducts()

	mockData := []viewModels.ProductListVM{
		{Id: 1, Name: "Qırmızı Alma", Quantity: 50, Weight: 120.5, BuyingPrice: 1.20, SellingPrice: 2.50, Stock: 50},
		{Id: 2, Name: "Sarı Armud", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Id: 3, Name: "Qara Tut", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Id: 4, Name: "Mavi Qizilgul", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Id: 5, Name: "Qirmizi Pomidor", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
	}

	buf := new(bytes.Buffer)
	product.List(mockData).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) ProductForm(id uint64) string {
	var p viewModels.ProductUpdateVM
	isEdit := id > 0

	if isEdit {
		// Normalda burada API-dan məhsulu ID-yə görə çəkməlisən
		// Məsələn: p, _ = a.Auth.API.GetProductByID(id)

		p = viewModels.ProductUpdateVM{
			Id:           id,
			Name:         "Sarı Armud",
			BuyingPrice:  2.10,
			SellingPrice: 3.80,
			Weight:       30.0,
			Stock:        5,
		}
	} else {
		p = viewModels.ProductUpdateVM{}
	}

	buf := new(bytes.Buffer)
	err := product.Form(p, isEdit).Render(context.Background(), buf)
	if err != nil {
		return "Form render xətası: " + err.Error()
	}

	return buf.String()
}
func (a *App) GetCustomerList() string {
	mockCustomers := []viewModels.CustomerResponseVM{
		{ID: 1, Name: "Əli", Surname: "Məmmədov", Phone: "055-123-45-67", Debt: 150.50},
		{ID: 2, Name: "Vəli", Surname: "Əliyev", Phone: "070-987-65-43", Debt: 0.00},
		{ID: 3, Name: "Vəli", Surname: "Əliyev", Phone: "070-987-65-43", Debt: 0.00},
		{ID: 4, Name: "Vəli", Surname: "Əliyev", Phone: "070-987-65-43", Debt: 0.00},
		{ID: 5, Name: "Vəli", Surname: "Əliyev", Phone: "070-987-65-43", Debt: 0.00},
	}
	buf := new(bytes.Buffer)
	customer.List(mockCustomers).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetCustomerForm(id uint64) string {
	var c viewModels.UpdateCustomerVM
	isEdit := id > 0
	if isEdit {
		c = viewModels.UpdateCustomerVM{ID: id, Name: "Əli", Surname: "Məmmədov", Phone: "055-123-45-67"}
	}
	buf := new(bytes.Buffer)
	customer.Form(c, isEdit).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetCustomerDetails(id uint64) string {
	c := viewModels.CustomerDetailsVM{
		ID:      id,
		Name:    "Albert",
		Surname: "Haciverdiyev",
		Phone:   "055-111-22-33",
		Address: "Baki",
		Debt:    54.50,
		Purchases: []viewModels.CustomerPurchaseVM{
			{ProductName: "Qırmızı Alma", Quantity: 5.5, Price: 2.00, TotalPrice: 11.00, Date: time.Now()},
			{ProductName: "Sarı Armud", Quantity: 2.0, Price: 3.50, TotalPrice: 7.00, Date: time.Now()},
		},
		Sales: []viewModels.CustomerSaleVM{
			{ProductName: "Qırmızı Alma", Quantity: 5.5, Price: 2.00, TotalPrice: 11.00, Date: time.Now()},
			{ProductName: "Sarı Armud", Quantity: 2.0, Price: 3.50, TotalPrice: 7.00, Date: time.Now()},
		},
		DebtLogs: []viewModels.CustomerDebtLogVM{
			{Amount: 18.00, Description: "Nisyə alış", Type: "increase", Date: time.Now()},
			{Amount: 10.00, Description: "Ödəniş edildi", Type: "decrease", Date: time.Now()},
		},
	}

	buf := new(bytes.Buffer)
	customer.Details(c).Render(context.Background(), buf)
	return buf.String()
}

// --- Sales ---

func (a *App) GetSaleList() string {
	// Realistik Satış Tarixçəsi Dump Datası
	mockSales := []viewModels.SaleVM{
		{
			Id:          101,
			Customer:    viewModels.CustomerResponseVM{Name: "Əli", Surname: "Məmmədov"},
			CrateDate:   time.Now().Add(-2 * time.Hour),
			Quantity:    15,
			Weight:      2.5,
			TotalWeight: 40.0, // Brutto
			NetWeight:   37.5, // Netto
			TotalPrice:  93.75,
		},
		{
			Id:          102,
			Customer:    viewModels.CustomerResponseVM{Name: "Albert", Surname: "Haciverdiyev"},
			CrateDate:   time.Now().Add(-24 * time.Hour),
			Quantity:    10,
			Weight:      1.8,
			TotalWeight: 20.0,
			NetWeight:   18.0,
			TotalPrice:  54.00,
		},
		{
			Id:          103,
			Customer:    viewModels.CustomerResponseVM{Name: "Nizami", Surname: "Gəncəvi"},
			CrateDate:   time.Now().Add(-48 * time.Hour),
			Quantity:    50,
			Weight:      0.5,
			TotalWeight: 28.0,
			NetWeight:   25.0,
			TotalPrice:  125.00,
		},
	}

	buf := new(bytes.Buffer)
	// Qeyd: views/sales paketindəki List funksiyasını çağırırıq
	sales.List(mockSales).Render(context.Background(), buf)
	// (Aşağıda birbaşa render üçün istifadə edə bilərsən)
	return buf.String()
}

func (a *App) GetSaleForm(id uint64) string {
	// Müştəri siyahısı (Select box üçün)
	mockCustomers := []viewModels.CustomerResponseVM{
		{ID: 1, Name: "Əli", Surname: "Məmmədov"},
		{ID: 3, Name: "Albert", Surname: "Haciverdiyev"},
		{ID: 4, Name: "Nizami", Surname: "Gəncəvi"},
	}

	var s viewModels.SaleRequestVM
	isEdit := id > 0

	if isEdit {
		// Redaktə rejimi üçün mövcud satış datası
		s = viewModels.SaleRequestVM{
			Id:          id,
			CustomerID:  3, // Albert
			Quantity:    12,
			Weight:      2.0,
			TotalWeight: 26.0,
			NetWeight:   24.0,
			TotalPrice:  48.00,
		}
	} else {
		// Yeni satış üçün boş model
		s = viewModels.SaleRequestVM{
			Quantity: 1,
			Weight:   0.0,
		}
	}

	buf := new(bytes.Buffer)
	sales.Form(mockCustomers, s, isEdit).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetPurchaseList() string {
	// Realistik Alış Tarixçəsi Dump Datası
	mockPurchases := []viewModels.PurchaseVM{
		{
			Id:          201,
			Customer:    viewModels.CustomerResponseVM{Name: "Tədarükçü", Surname: "Xəzər MMC"},
			CrateDate:   time.Now().Add(-5 * time.Hour),
			Quantity:    100,
			Weight:      0.8,
			TotalWeight: 85.0, // Brutto
			NetWeight:   80.0, // Netto
			TotalPrice:  320.00,
		},
		{
			Id:          202,
			Customer:    viewModels.CustomerResponseVM{Name: "Vüqar", Surname: "Qasımov"},
			CrateDate:   time.Now().Add(-48 * time.Hour),
			Quantity:    50,
			Weight:      1.2,
			TotalWeight: 65.0,
			NetWeight:   60.0,
			TotalPrice:  180.00,
		},
		{
			Id:          203,
			Customer:    viewModels.CustomerResponseVM{Name: "Böyük", Surname: "Anbar Şirkəti"},
			CrateDate:   time.Now().Add(-120 * time.Hour),
			Quantity:    200,
			Weight:      0.5,
			TotalWeight: 110.0,
			NetWeight:   100.0,
			TotalPrice:  450.00,
		},
	}

	buf := new(bytes.Buffer)
	// Qeyd: views/purchases paketindəki List funksiyasını çağırırıq
	purchases.List(mockPurchases).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetPurchaseForm(id uint64) string {
	// Tədarükçü/Müştəri siyahısı (Select box üçün)
	mockSuppliers := []viewModels.CustomerResponseVM{
		{ID: 10, Name: "Tədarükçü", Surname: "Xəzər MMC"},
		{ID: 11, Name: "Vüqar", Surname: "Qasımov"},
		{ID: 12, Name: "Böyük", Surname: "Anbar Şirkəti"},
	}

	var p viewModels.PurchaseRequestVM
	isEdit := id > 0

	if isEdit {
		// Redaktə rejimi üçün mövcud alış datası (Dump)
		p = viewModels.PurchaseRequestVM{
			Id:          id,
			CustomerID:  11, // Vüqar Qasımov
			Quantity:    50,
			Weight:      1.2,
			TotalWeight: 65.0,
			NetWeight:   60.0,
			TotalPrice:  180.00,
		}
	} else {
		// Yeni alış üçün standart boş model
		p = viewModels.PurchaseRequestVM{
			Quantity: 0,
			Weight:   0.0,
		}
	}

	buf := new(bytes.Buffer)
	// views/purchases paketindəki Form funksiyasını çağırırıq
	purchases.Form(mockSuppliers, p, isEdit).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) SetToken(token string) {
	a.Auth.API.Token = token
	fmt.Println("Köhnə token bərpa edildi.")
}

func (a *App) GetDashboard() string {
	// Digər metodlardakı mock dataları birləşdirək
	data := viewModels.DashboardVM{
		TotalSalesToday:     1250.40,
		TotalPurchasesToday: 840.20,
		TotalStockWeight:    4280.50,
		ActiveCustomerCount: 142,
		RecentTransactions: []viewModels.TransactionVM{
			{Type: "sale", Party: "Əli Məmmədov", Amount: 45.50, Time: "14:20"},
			{Type: "purchase", Party: "Xəzər MMC", Amount: 120.00, Time: "12:15"},
			{Type: "sale", Party: "Albert Haciverdiyev", Amount: 12.20, Time: "10:05"},
		},
		LowStockProducts: []viewModels.ProductListVM{
			{Name: "Sarı Armud", Weight: 5.0},
			{Name: "Qırmızı Alma", Weight: 8.5},
		},
	}

	buf := new(bytes.Buffer)
	dashboard.List(data).Render(context.Background(), buf)
	return buf.String()
}
