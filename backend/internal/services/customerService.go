package services

import (
	"context"
	"fmt"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/dto"
	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) GetAll(ctx context.Context, search string) ([]entity.Customer, error) {
	var customers []entity.Customer
	query := s.db.WithContext(ctx).Model(&entity.Customer{})

	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		query = query.Where("name LIKE ? OR surname LIKE ? OR email LIKE ?", searchTerm, searchTerm, searchTerm)
	}

	err := query.Find(&customers).Error
	return customers, err
}

func (s *CustomerService) GetByID(ctx context.Context, id uint64) (*entity.Customer, error) {
	var customer entity.Customer
	err := s.db.WithContext(ctx).First(&customer, id).Error
	return &customer, err
}

func (s *CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := entity.Customer{
		Name:    req.Name,
		Surname: req.Surname,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Image:   req.Image,
	}
	return s.db.WithContext(ctx).Create(&customer).Error
}

func (s *CustomerService) Update(ctx context.Context, id uint64, req dto.UpdateCustomerRequest) error {
	updates := map[string]interface{}{
		"name":    req.Name,
		"surname": req.Surname,
		"email":   req.Email,
		"phone":   req.Phone,
		"address": req.Address,
		"image":   req.Image,
	}
	return s.db.WithContext(ctx).Model(&entity.Customer{}).Where("id = ?", id).Updates(updates).Error
}

func (s *CustomerService) Delete(ctx context.Context, id uint64) error {
	return s.db.WithContext(ctx).Delete(&entity.Customer{}, id).Error
}
