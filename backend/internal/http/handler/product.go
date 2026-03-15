package handler

import (
	"net/http"
	"strconv"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/dto"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) List(c *gin.Context) {
	searchQuery := c.Query("q")

	p, err := h.service.GetAll(c.Request.Context(), searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product List Error"})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product ID Error"})
		return
	}
	p, err := h.service.GetByID(c.Request.Context(), ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product GetByID Error"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Məlumatlar yanlışdır",
			"details": err.Error(),
		})
		return
	}

	product := entity.Product{
		Name:         req.Name,
		Quantity:     req.Quantity,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Weight:       req.Weight,
		Stock:        req.Stock,
	}

	if err := h.service.Create(c.Request.Context(), product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Məhsul yaradılarkən xata baş verdi"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Məhsul uğurla yaradıldı"})
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kecersiz ID formatı"})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz güncelleme verisi"})
		return
	}

	product := entity.Product{
		Name:         req.Name,
		Quantity:     req.Quantity,
		BuyingPrice:  req.BuyingPrice,
		SellingPrice: req.SellingPrice,
		Weight:       req.Weight,
		Stock:        req.Stock,
	}
	product.ID = uint(id)

	if err := h.service.Update(c.Request.Context(), product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product Update Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Update Success"})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product Delete Error"})
		return
	}
	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product Delete Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Delete Success"})
}
