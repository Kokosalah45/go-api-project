package repository

import "go-api-project/internal/features/users/model"

type UserSchema struct {
	ID          int     `bson:"_id,omitempty" json:"id"`
	Username    string  `bson:"username" json:"username"`
	Email       string  `bson:"email" json:"email"`
	Password    string  `bson:"password" json:"-"`
	IsAdmin     bool    `bson:"is_admin" json:"is_admin"`
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
}

func (s *UserSchema) ToModel() *model.User {
	return &model.User{
		ID:          s.ID,
		Username:    s.Username,
		Email:       s.Email,
		IsAdmin:     s.IsAdmin,
		Description: s.Description,
	}
}

func FromModel(m *model.User) *UserSchema {
	return &UserSchema{
		ID:          m.ID,
		Username:    m.Username,
		Email:       m.Email,
		IsAdmin:     m.IsAdmin,
		Description: m.Description,
	}
}
