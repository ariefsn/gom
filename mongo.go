package gom

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo = Mongo struct
type Mongo struct {
	Config           MongoConfig
	Client           *mongo.Client
	Context          context.Context
	Collection       *mongo.Collection
	ContextCancel    context.CancelFunc
	ConnectionString string
}

// MongoConfig = Mongo config struct
type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// NewMongo = Init new mongo
func NewMongo() *Mongo {
	return new(Mongo)
}

// SetContext = Set context with seconds as param
func (m *Mongo) SetContext(seconds time.Duration) {
	if &seconds == nil {
		seconds = 30
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), seconds*time.Second)

	m.Context = ctx
	m.ContextCancel = ctxCancel
}

// SetConfig = Set config of mongo
func (m *Mongo) SetConfig(config MongoConfig) {
	m.Config = config
}

// SetClient = Set client
func (m *Mongo) SetClient() {
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
