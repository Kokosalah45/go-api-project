package service

import (
	"context"
	"errors"
	"go-api-project/internal/features/users/domain"
)

type UserServicer interface {
	CreateUser(ctx context.Context, createUserDTO *domain.User) (int, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, errors.New("failed to create user: " + err.Error())
	}
	return userID, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}
	return user, nil
}
