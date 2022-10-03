package interfaces

import (
	"context"
	"edward/cleanarc/pkg/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id uint) (domain.User, error)
}