package ports

import (
	"context"
	"go-api-project/internal/features/users/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (int, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
}
