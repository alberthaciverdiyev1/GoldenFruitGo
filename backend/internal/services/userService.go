package services

import (
	"context"
	"errors"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/entity"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	jwt *JWTService
}

func NewUserService(db *gorm.DB, jwt *JWTService) *UserService {
	return &UserService{db: db, jwt: jwt}
}

func (u *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (*dto.UserResponse, error) {
	var user entity.User

	err := u.db.WithContext(ctx).Where("user_name = ?", req.UserName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("istifadeci tapilmadi")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("sifre sehvdir")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return nil, errors.New("cant generate token")
	}

	return &dto.UserResponse{
		UserID:   user.ID,
		UserName: user.UserName,
		Token:    token,
	}, nil
}

func (u *UserService) Logout(ctx context.Context) error {
	return nil
}
