package repository

import (
	"context"
	"fmt"
	"go-api-project/internal/database"
	mongodb "go-api-project/internal/database/mongodb"
	"go-api-project/bff-users/features/users/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const collectionName = "users"

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db database.DBService) *MongoUserRepository {
	mongoDB, ok := db.(*mongodb.MongoDbService)
	if !ok {
		panic("Failed to assert database client as *mongo.Client")
	}
	return &MongoUserRepository{
		collection: mongoDB.GetDb().Collection(collectionName),
	}
}

func (m *MongoUserRepository) Create(ctx context.Context, user *domain.User) (int, error) {

	fmt.Print(user)

	userEntity, mappingErr := FromModel(user)

	if mappingErr != nil {
		return 0, mappingErr
	}

	collection := m.collection

	_, err := collection.InsertOne(ctx, userEntity)

	fmt.Print(err)

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (m *MongoUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {

	var userEntity UserSchema

	collection := m.collection

	parsedId, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	result := collection.FindOne(ctx, bson.M{"_id": parsedId})

	if err := result.Decode(&userEntity); err != nil {
		return nil, err
	}

	if userEntity.ID.IsZero() {
		return nil, fmt.Errorf("user not found")
	}

	return userEntity.ToModel(), nil
}

func (m *MongoUserRepository) Update(ctx context.Context, user *domain.User) error {

	collection := m.collection

	userEntity, mappingErr := FromModel(user)

	if mappingErr != nil {
		return mappingErr
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": userEntity.ID}, bson.M{"$set": userEntity})

	if err != nil {
		return err
	}

	return nil
}
