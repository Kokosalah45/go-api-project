package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


type MongoDbService struct {
	client *mongo.Client
	dbName string
	db     *mongo.Database
}

type MongoDBConf struct {
	Host     string
	Port     string
	AppUser  string
	AppPass  string
	DBName   string
}


func New(c *MongoDBConf) (*MongoDbService, error) {

	appURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=%s",
		c.AppUser, c.AppPass, c.Host, c.Port, c.DBName, c.DBName,
	)
	appClient, err := mongo.Connect(options.Client().ApplyURI(appURI))
	
	if err != nil {
		return nil, fmt.Errorf("connecting to MongoDB: %w", err)
	}
	svc := &MongoDbService{
		client: appClient,
		dbName: c.DBName,
		db:     appClient.Database(c.DBName),
	}

	return svc, nil
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
