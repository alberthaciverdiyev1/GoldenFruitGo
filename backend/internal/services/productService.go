package services

import (
	"context"
	"fmt"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{DB: db}
}

func (s *ProductService) GetAll(ctx context.Context, search string) ([]entity.Product, error) {
	products := []entity.Product{}

	query := s.DB.WithContext(ctx).Model(&entity.Product{})

	if search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", search)
		query = query.Where("name LIKE ? ", searchTerm)
	}

	err := query.Order("id DESC").Find(&products).Error

	return products, err
}

func (s *ProductService) GetByID(ctx context.Context, id uint64) (entity.Product, error) {
	product := entity.Product{}
	err := s.DB.First(&product, id).Error
	return product, err
}

func (s *ProductService) Create(ctx context.Context, product entity.Product) error {

	err := s.DB.WithContext(ctx).Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Update(ctx context.Context, product entity.Product) error {

	err := s.DB.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", product.ID).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) Delete(ctx context.Context, id uint64) error {
	err := s.DB.WithContext(ctx).Delete(&entity.Product{}, id).Error
	return err
}
