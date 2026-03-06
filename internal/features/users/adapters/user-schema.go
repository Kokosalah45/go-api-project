package adapters

import (
	"go-api-project/internal/features/users/domain"

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
	}
}

func FromModel(m *domain.User) *UserSchema {
	parsedId, err := bson.ObjectIDFromHex(m.ID)

	if err != nil {
		parsedId = bson.NewObjectID()
	}

	return &UserSchema{
		ID:          parsedId,
		Username:    m.Username,
		Email:       m.Email,
		Description: m.Description,
	}
}
