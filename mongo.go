package gom

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// ScramSha1 is SCRAM-SHA-1 (legacy)
	ScramSha1 = "SCRAM-SHA-1"
	// ScramSha256 is SCRAM-SHA-256
	ScramSha256 = "SCRAM-SHA-256"
	// MongoDbCr is MONGODB-CR
	MongoDbCr = "MONGODB-CR"
	// Plain is PLAIN
	Plain = "PLAIN"
	// GssAPI is GSSAPI
	GssAPI = "GSSAPI"
	// MongoDbX509 is MONGODB-X509
	MongoDbX509 = "MONGODB-X509"
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
	Username      string
	Password      string
	Host          string
	Port          int
	Database      string
	MaxPool       int
	AuthMechanism string
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

	connectionString := fmt.Sprintf("mongodb://%s:%v", config.Host, config.Port)

	if config.Username != "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%v", config.Username, config.Password, config.Host, config.Port)
		if config.AuthMechanism != "" {
			connectionString = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?authMechanism=%s", config.Username, config.Password, config.Host, config.Database, config.AuthMechanism)
		}
	}

	clientOptions := options.Client().ApplyURI(connectionString)

	if config.MaxPool > 0 {
		clientOptions.SetMaxPoolSize(uint64(config.MaxPool))
	}

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
