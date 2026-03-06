package domain

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (int, error)
	GetByID(ctx context.Context, id string) (*User, error)
}
