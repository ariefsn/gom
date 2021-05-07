package gom

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Gom struct
type Gom struct {
	mongo mongoDB
}

// NewGom = Create new
func NewGom() *Gom {
	return new(Gom)
}

// Init = Init
func (g *Gom) Init(config Config) {
	g.mongo.SetConfig(config)
	g.mongo.SetClient()
}

// Set = Get set query with gom
func (g *Gom) Set(SetParams *SetParams) *Set {
	s := newSet(g, SetParams)

	return s
}

// CheckClient = Check connection successfull or not
func (g *Gom) CheckClient() error {
	err := g.mongo.Client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return errors.New(toolkit.Sprintf("Couldn't connect to database : %s", err.Error()))
	}

	toolkit.Println(toolkit.Sprintf("Connected to database: %s", g.mongo.ConnectionString))

	return nil
}

// GetClient = Get active client
func (g *Gom) GetClient() *mongo.Client {
	return g.mongo.Client
}

// GetDatabase = Get database name
func (g *Gom) GetDatabase() string {
	return g.mongo.Config.Database
}
