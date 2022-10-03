package repository

import (
	"context"
	"edward/cleanarc/pkg/domain"
	interfaces "edward/cleanarc/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

// FindAll implements interfaces.UserRepository
func (c *userDatabase) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := c.DB.Find(&users).Error
	//panic("unimplemented")

	return users, err
}

// FindById implements interfaces.UserRepository
func (c *userDatabase) FindById(ctx context.Context, id uint) (domain.User, error) {
	//panic("unimplemented")
	var user domain.User
	err := c.DB.First(&user, id).Error

	return user, err
}
