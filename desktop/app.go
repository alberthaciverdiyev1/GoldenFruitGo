package main

import (
	"bytes"
	"context"
	"desktop/internal/api"
	"desktop/internal/handlers"
	"desktop/internal/viewModels"
	"desktop/internal/views/product"
	"fmt"
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
		{Name: "Qırmızı Alma", Quantity: 50, Weight: 120.5, BuyingPrice: 1.20, SellingPrice: 2.50, Stock: 50},
		{Name: "Sarı Armud", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Name: "Qara Tut", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Name: "Mavi Qizilgul", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
		{Name: "Qirmizi Pomidor", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
	}

	buf := new(bytes.Buffer)
	product.List(mockData).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) SetToken(token string) {
	a.Auth.API.Token = token
	fmt.Println("Köhnə token bərpa edildi.")
}
