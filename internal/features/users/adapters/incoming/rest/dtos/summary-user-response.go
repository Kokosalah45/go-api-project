package dtos

import "go-api-project/internal/features/users/domain"

type SummaryUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewSummaryUserResponse(user *domain.User) *SummaryUserResponse {
	return &SummaryUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
