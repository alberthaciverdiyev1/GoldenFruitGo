package handler

import (
	"bytes"
	"context"
	"desktop/internal/viewModels"
	sales "desktop/internal/views/sale"
	"time"
)

func (a *App) GetSaleList() string {
	mockSales := []viewModels.SaleVM{
		{Id: 101, Customer: viewModels.CustomerResponseVM{Name: "Əli", Surname: "Məmmədov"}, CrateDate: time.Now(), TotalPrice: 93.75},
	}
	buf := new(bytes.Buffer)
	sales.List(mockSales).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetSaleForm(id uint64) string {
	mockCustomers := []viewModels.CustomerResponseVM{{ID: 1, Name: "Əli", Surname: "Məmmədov"}}
	var s viewModels.SaleRequestVM
	isEdit := id > 0
	buf := new(bytes.Buffer)
	sales.Form(mockCustomers, s, isEdit).Render(context.Background(), buf)
	return buf.String()
}
