package repository

import (
	"go-api-project/bff-users/features/users/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserSchema struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Username    string        `bson:"username" `
	Email       string        `bson:"email"`
	Description *string       `bson:"description,omitempty"`
	Age         *int          `bson:"age,omitempty"`
}

func (s *UserSchema) ToModel() *domain.User {

	return &domain.User{
		ID:          s.ID.Hex(),
		Username:    s.Username,
		Email:       s.Email,
		Description: s.Description,
		Age:         s.Age,
	}
}

func FromModel(m *domain.User) (*UserSchema, error) {
	parsedId, err := bson.ObjectIDFromHex(m.ID)

	if err != nil {
		return nil, err
	}

	return &UserSchema{
		ID:          parsedId,
		Username:    m.Username,
		Email:       m.Email,
		Description: m.Description,
		Age:         m.Age,
	} , nil
}
