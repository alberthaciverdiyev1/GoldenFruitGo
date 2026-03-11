package handler

import (
	"bytes"
	"context"
	"desktop/internal/api"
	"desktop/internal/viewModels"
	"desktop/internal/views/dashboard"
	"fmt"
)

type App struct {
	ctx context.Context
	API *api.Client
}

func NewApp(apiUrl string) *App {
	return &App{
		API: api.NewClient(apiUrl),
	}
}

func (a *App) Startup(ctx context.Context) { a.ctx = ctx }

func (a *App) SetToken(token string) {
	a.API.Token = token
	fmt.Println("Token bərpa edildi.")
}

func (a *App) GetDashboard() string {
	data := viewModels.DashboardVM{
		TotalSalesToday:     1250.40,
		TotalPurchasesToday: 840.20,
		TotalStockWeight:    4280.50,
		ActiveCustomerCount: 142,
		RecentTransactions: []viewModels.TransactionVM{
			{Type: "sale", Party: "Əli Məmmədov", Amount: 45.50, Time: "14:20"},
			{Type: "purchase", Party: "Xəzər MMC", Amount: 120.00, Time: "12:15"},
		},
	}
	buf := new(bytes.Buffer)
	dashboard.List(data).Render(context.Background(), buf)
	return buf.String()
}
