package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbService struct {
	client *mongo.Client
	dbName string
	db    *mongo.Database
}

const databaseName = "brokers"

var (
	host = os.Getenv("BLUEPRINT_DB_HOST")
	port = os.Getenv("BLUEPRINT_DB_PORT")
)

func New() *MongoDbService {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port)),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoDbService{
		client: client,
		dbName: databaseName,
		db:    client.Database(databaseName),

	}
}

func (s *MongoDbService) GetClient() *mongo.Client {
	if s.client == nil {
		log.Fatal("Database connection is not initialized")
	}
	return s.client
}


func (s *MongoDbService) GetDBName() string {
	return s.dbName
}

func (s *MongoDbService) GetDb() *mongo.Database {
	if s.db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return s.db
}


func (s *MongoDbService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *MongoDbService) Close() error {
	if s.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return s.client.Disconnect(context.Background())
}