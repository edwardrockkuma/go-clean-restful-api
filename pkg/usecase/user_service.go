package usecase

import (
	"context"
	"edward/cleanarc/pkg/domain"
	interfaces "edward/cleanarc/pkg/repository/interface"
	services "edward/cleanarc/pkg/usecase/interface"
)

type userService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) services.UserService {
	return &userService{
		userRepo: repo,
	}
}

// Delete implements interfaces.UserService
func (c *userService) Delete(ctx context.Context, user domain.User) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}

// FindAll implements interfaces.UserService
func (c *userService) FindAll(ctx context.Context) ([]domain.User, error) {
	users, err := c.userRepo.FindAll(ctx)

	return users, err
}

// FindByID implements interfaces.UserService
func (c *userService) FindById(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.userRepo.FindById(ctx, id)

	return user, err
}

// Save implements interfaces.UserService
func (c *userService) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}
