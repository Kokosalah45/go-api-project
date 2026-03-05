package service

import (
	"context"
	"errors"
	"go-api-project/internal/features/users/model"
	"go-api-project/internal/features/users/repository"
)

type UserServicer interface{
	CreateUser(ctx context.Context, createUserDTO *model.User) (int, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (int, error) {
	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, errors.New("failed to create user: " + err.Error())
	}
	return userID, nil
}


func (s *UserService) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}
	return user, nil
}