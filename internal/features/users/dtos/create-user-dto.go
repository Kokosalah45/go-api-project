package dtos

import "go-api-project/internal/features/users/model"

type CreateUserDTO struct {
	Username    string  `json:"username" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`
	IsAdmin     bool    `json:"is_admin"`
	Description *string `json:"description"`
}

func (dto *CreateUserDTO) ToModel() *model.User {
	return &model.User{
		Username:    dto.Username,
		Email:       dto.Email,
		IsAdmin:     dto.IsAdmin,
		Description: dto.Description,
	}
}
