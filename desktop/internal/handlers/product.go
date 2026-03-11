package handler

import (
	"bytes"
	"context"
	"desktop/internal/viewModels"
	"desktop/internal/views/product"
)

func (a *App) GetProductListHTML() string {
	// Gələcəkdə: products, _ := a.API.Get("/products")
	mockData := []viewModels.ProductListVM{
		{Id: 1, Name: "Qırmızı Alma", Quantity: 50, Weight: 120.5, BuyingPrice: 1.20, SellingPrice: 2.50, Stock: 50},
		{Id: 2, Name: "Sarı Armud", Quantity: 12, Weight: 30.0, BuyingPrice: 2.10, SellingPrice: 3.80, Stock: 5},
	}

	buf := new(bytes.Buffer)
	product.List(mockData).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) ProductForm(id uint64) string {
	var p viewModels.ProductUpdateVM
	isEdit := id > 0
	if isEdit {
		p = viewModels.ProductUpdateVM{Id: id, Name: "Sarı Armud", BuyingPrice: 2.10, SellingPrice: 3.80, Weight: 30.0, Stock: 5}
	}
	buf := new(bytes.Buffer)
	product.Form(p, isEdit).Render(context.Background(), buf)
	return buf.String()
}
