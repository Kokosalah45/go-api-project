package service

import (
	"context"
	"errors"
	"go-api-project/internal/features/users/adapters/incoming/rest/dtos"
	"go-api-project/internal/features/users/domain"
	"go-api-project/internal/features/users/ports"
)



type UserServicer interface {
	CreateUser(ctx context.Context, createUserDTO *domain.User) (int, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
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

func (s *UserService) UpdateUser(ctx context.Context, id string, updatedUser *dtos.UpdateUserRequest) (*domain.User, error) {

	existingUser, err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil, errors.New("failed to get user: " + err.Error())
	}

	if updatedUser.Username != nil {
		existingUser.Username = *updatedUser.Username
	}

	if updatedUser.Email != nil {
		existingUser.Email = *updatedUser.Email
	}

	if updatedUser.Description != nil {
		existingUser.Description = updatedUser.Description
	}

	if updatedUser.Age != nil {
		existingUser.Age = updatedUser.Age
	}

	err = s.repo.Update(ctx, existingUser)

	if err != nil {
		return nil, errors.New("failed to update user: " + err.Error())
	}

	return existingUser, nil
}
