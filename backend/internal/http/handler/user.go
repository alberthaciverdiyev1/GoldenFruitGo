package handler

import (
	"net/http"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/dto"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Məlumatlar düzgün deyil",
			"details": err.Error(),
		})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Logout - Gin uyumlu handler
func (h *UserHandler) Logout(c *gin.Context) {
	err := h.service.Logout(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sistemden çıxış zamanı xəta baş verdi",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Uğurla çıxış edildi",
	})
}
