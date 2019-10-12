package gom

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoDB = MongoDB struct
type mongoDB struct {
	Config           Config
	Client           *mongo.Client
	Collection       *mongo.Collection
	ConnectionString string
}

// Config = Mongo config struct
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// newMongo = Init new mongo
func newMongo() *mongoDB {
	return new(mongoDB)
}

// SetConfig = Set config of mongo
func (m *mongoDB) SetConfig(config Config) {
	m.Config = config
}

// SetClient = Set client
func (m *mongoDB) SetClient() {
	config := m.Config

	connectionString := fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)

	if config.Username != "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Username, config.Password, config.Host, config.Port)
	}

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	m.ConnectionString = connectionString

	m.Client = client
}
