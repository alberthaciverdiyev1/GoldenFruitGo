package main

import (
	"context"
	"desktop/internal/api"
	"desktop/internal/handlers"
	"desktop/internal/viewModels"
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

//func (a *App) DoLogin(user, pass string) string {
//	fmt.Println("do login")
//	res, err := a.Auth.Login(user, pass)
//	if err != nil {
//		return err.Error()
//	}
//	return "success:" + res.UserName
//}

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
