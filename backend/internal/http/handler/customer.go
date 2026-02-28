package handler

import (
	"net/http"
	"strconv"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/dto"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/services"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// List: GET /api/v1/customers?q=search_term
func (h *CustomerHandler) List(c *gin.Context) {
	searchQuery := c.Query("q")

	customers, err := h.service.GetAll(c.Request.Context(), searchQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Müşteriler listelenemedi"})
		return
	}

	// Liste boş olsa bile 200 OK ve boş array dönmek best practice'dir.
	c.JSON(http.StatusOK, customers)
}

// GetByID: GET /api/v1/customers/:id
func (h *CustomerHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID formatı"})
		return
	}

	customer, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		// Kayıt bulunamadığında 404 dönüyoruz
		c.JSON(http.StatusNotFound, gin.H{"error": "Müşteri bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Create: POST /api/v1/customers
func (h *CustomerHandler) Create(c *gin.Context) {
	var req dto.CreateCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri", "details": err.Error()})
		return
	}

	if err := h.service.Create(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Müşteri oluşturulurken bir hata oluştu"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Müşteri başarıyla oluşturuldu"})
}

// Update: PUT /api/v1/customers/:id
func (h *CustomerHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID formatı"})
		return
	}

	var req dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz güncelleme verisi"})
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Müşteri güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Müşteri başarıyla güncellendi"})
}

// Delete: DELETE /api/v1/customers/:id
func (h *CustomerHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID formatı"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Silme işlemi başarısız"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Müşteri başarıyla silindi"})
}
