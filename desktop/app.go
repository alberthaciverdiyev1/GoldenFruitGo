package main

import (
	"bytes"
	"context"
	"desktop/internal/api"
	"desktop/internal/handlers"
	"desktop/internal/viewModels"
	"desktop/internal/views/customer"
	"desktop/internal/views/product"
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

func (a *App) SetToken(token string) {
	a.Auth.API.Token = token
	fmt.Println("Köhnə token bərpa edildi.")
}
