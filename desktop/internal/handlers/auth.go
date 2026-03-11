package handler

import (
	"bytes"
	"context"
	"desktop/internal/viewModels"
	auth "desktop/internal/views/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) GetLoginPageHTML(errMsg string) string {
	buf := new(bytes.Buffer)
	auth.LoginPage(errMsg).Render(context.Background(), buf)
	return buf.String()
}

func (a *App) DoLogin(userName, password string) (*viewModels.UserResponse, error) {
	req := viewModels.UserLoginRequest{
		UserName: userName,
		Password: password,
	}

	resp, err := a.API.Post("/auth/login", req)
	if err != nil {
		return nil, fmt.Errorf("API bağlantı xətası: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("İstifadəçi adı və ya şifrə yanlışdır")
	}

	var res viewModels.UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	a.API.Token = res.Token
	return &res, nil
}
