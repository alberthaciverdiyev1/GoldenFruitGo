package handler

import (
	"bytes"
	"context"
	"desktop/internal/viewModels"
	purchases "desktop/internal/views/purchase"
	"time"
)

func (a *App) GetPurchaseList() string {
	mockPurchases := []viewModels.PurchaseVM{
		{Id: 201, Customer: viewModels.CustomerResponseVM{Name: "Tədarükçü", Surname: "Xəzər MMC"}, CrateDate: time.Now(), TotalPrice: 320.00},
	}
	buf := new(bytes.Buffer)
	purchases.List(mockPurchases).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) GetPurchaseForm(id uint64) string {
	mockSuppliers := []viewModels.CustomerResponseVM{{ID: 10, Name: "Tədarükçü", Surname: "Xəzər MMC"}}
	var p viewModels.PurchaseRequestVM
	isEdit := id > 0
	buf := new(bytes.Buffer)
	purchases.Form(mockSuppliers, p, isEdit).Render(context.Background(), buf)
	return buf.String()
}
