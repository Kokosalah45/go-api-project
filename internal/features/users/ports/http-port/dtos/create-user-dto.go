package dtos

import (
	"go-api-project/internal/features/users/domain"
)

type CreateUserDTO struct {
	Username    string  `json:"username" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Description *string `json:"description" binding:"omitempty,min=25,max=500"`
	Age         *int    `json:"age" binding:"omitempty,gte=25,lte=55"`
}

func (dto *CreateUserDTO) ToModel() *domain.User {

	return &domain.User{
		Username:    dto.Username,
		Email:       dto.Email,
		Description: dto.Description,
		Age:         dto.Age,
	}
}
