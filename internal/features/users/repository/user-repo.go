package repository

import (
	"context"
	"go-api-project/internal/features/users/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}
