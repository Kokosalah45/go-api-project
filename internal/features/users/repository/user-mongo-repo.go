package repository

import (
	"context"
	"go-api-project/internal/database"
	"go-api-project/internal/features/users/model"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "users"

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db database.DBService) *MongoUserRepository {
	mongoDB, ok := db.(*database.MongoDbService)
	if !ok {
		panic("Failed to assert database client as *mongo.Client")
	}
	return &MongoUserRepository{
		collection: mongoDB.GetDb().Collection(collectionName),
	}
}

func (m *MongoUserRepository) Create(ctx context.Context, user *model.User) (int, error) {

	userEntity := FromModel(user)

	collection := m.collection

	_, err := collection.InsertOne(ctx, userEntity)

	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (m *MongoUserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {

	var userEntity UserSchema

	collection := m.collection

	err := collection.FindOne(ctx, map[string]interface{}{"_id": id}).Decode(&userEntity)
	
	if err != nil {
		return nil, err
	}

	return userEntity.ToModel(), nil
}
