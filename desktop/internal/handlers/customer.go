package handler

import (
	"bytes"
	"context"
	"desktop/internal/viewModels"
	"desktop/internal/views/customer"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func (a *App) GetCustomerList(search string) string {
	endpoint := "/customers?search=" + url.QueryEscape(search)
	resp, err := a.API.Get(endpoint)
	if err != nil {
		return "Xəta: " + err.Error()
	}
	defer resp.Body.Close()

	// 2. Decode et
	var customers []viewModels.CustomerResponseVM
	json.NewDecoder(resp.Body).Decode(&customers)

	// 3. Render et və HTML string qaytar
	buf := new(bytes.Buffer)
	customer.List(customers).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) CreateCustomer(data viewModels.UpdateCustomerVM) string {

	resp, err := a.API.Post("/customers", data)
	if err != nil {
		return "Xəta: " + err.Error()
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	responseString := string(bodyBytes)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Sprintf("Server xətası (%d): %s", resp.StatusCode, responseString)
	}
	return "Uğurlu"
}

// GetCustomerForm - Formu render edir
func (a *App) GetCustomerForm(id uint64) string {
	var c viewModels.UpdateCustomerVM
	isEdit := id > 0

	if isEdit {
		resp, err := a.API.Get(fmt.Sprintf("/customers/%d", id))
		if err == nil {
			defer resp.Body.Close()
			var res viewModels.CustomerResponseVM
			json.NewDecoder(resp.Body).Decode(&res)
			c = viewModels.UpdateCustomerVM{
				ID: res.ID, Name: res.Name, Surname: res.Surname, Phone: res.Phone, Address: res.Address,
			}
		}
	}

	buf := new(bytes.Buffer)
	customer.Form(c, isEdit).Render(context.Background(), buf)
	return buf.String()
}

// DeleteCustomer
func (a *App) DeleteCustomer(id uint64) string {
	resp, err := a.API.Delete(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return "Xəta: " + err.Error()
	}
	defer resp.Body.Close()
	return "Uğurlu"
}

// GetCustomerDetails - Müştərinin ətraflı məlumat səhifəsini render edir
func (a *App) GetCustomerDetails(id uint64) string {
	// 1. API-dan datanı çək
	resp, err := a.API.Get(fmt.Sprintf("/customers/%d", id))
	if err != nil {
		return "Xəta: " + err.Error()
	}
	defer resp.Body.Close()

	var res viewModels.CustomerDetailsVM
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "Decode xətası: " + err.Error()
	}

	// 2. Render et (customer view-dakı Details funksiyası ilə)
	buf := new(bytes.Buffer)
	customer.Details(res).Render(context.Background(), buf)
	return buf.String()
}

// UpdateCustomer - Mövcud müştərini yeniləyir
func (a *App) UpdateCustomer(data viewModels.UpdateCustomerVM) string {
	// API-da UPDATE (PUT) metodunu çağırırıq
	resp, err := a.API.Put(fmt.Sprintf("/customers/%d", data.ID), data)
	if err != nil {
		return "Xəta: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Sprintf("Server xətası: %d", resp.StatusCode)
	}
	return "Uğurlu"
}
