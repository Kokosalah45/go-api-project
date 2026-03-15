package dtos

import "go-api-project/bff-users/features/users/domain"

type DetailedUserResponse struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Description *string `json:"description,omitempty"`
	Age         *int    `json:"age,omitempty"`
}

func NewDetailedUserResponse(user *domain.User) *DetailedUserResponse {
	return &DetailedUserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Description: user.Description,
		Age:         user.Age,
	}
}
