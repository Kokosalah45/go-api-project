package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	host    = os.Getenv("DB_HOST")
	port    = os.Getenv("DB_PORT")
	appUser = os.Getenv("DB_USERNAME")
	appPass = os.Getenv("DB_PASSWORD")
	dbName  = os.Getenv("DB_NAME")
)

type MongoDbService struct {
	client *mongo.Client
	dbName string
	db     *mongo.Database
}

func New() *MongoDbService {

	appURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		appUser, appPass, host, port, dbName, dbName,
	)
	appClient, err := mongo.Connect(options.Client().ApplyURI(appURI))
	if err != nil {
		log.Fatalf("Failed to connect as app user: %v", err)
	}
	svc := &MongoDbService{
		client: appClient,
		dbName: dbName,
		db:     appClient.Database(dbName),
	}

	return svc
}

func (s *MongoDbService) GetDb() *mongo.Database   { return s.db }
func (s *MongoDbService) GetClient() *mongo.Client { return s.client }
func (s *MongoDbService) GetDBName() string        { return s.dbName }

func (s *MongoDbService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.client.Ping(ctx, nil); err != nil {
		log.Fatalf("db down: %v", err)
	}
	return map[string]string{"message": "It's healthy"}
}

func (s *MongoDbService) Close() error {
	return s.client.Disconnect(context.Background())
}
