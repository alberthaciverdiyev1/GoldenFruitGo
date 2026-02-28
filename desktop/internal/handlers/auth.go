package handlers

import (
	"bytes"
	"context"
	"desktop/internal/api"
	"desktop/internal/viewModels"
	auth "desktop/internal/views/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	API *api.Client
}

func (h *AuthHandler) GetLoginPageHTML(errMsg string) string {
	buf := new(bytes.Buffer)
	auth.LoginPage(errMsg).Render(context.Background(), buf)
	return buf.String()
}

func (h *AuthHandler) Login(userName, password string) (*viewModels.UserResponse, error) {
	req := viewModels.UserLoginRequest{
		UserName: userName,
		Password: password,
	}

	resp, err := h.API.Post("/auth/login", req)
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

	h.API.Token = res.Token
	return &res, nil
}
