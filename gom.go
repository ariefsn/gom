package gom

import (
	"context"
	"errors"
	"time"

	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Gom struct
type Gom struct {
	Mongo Mongo
}

// NewGom = Create new
func NewGom() *Gom {
	return new(Gom)
}

// Init = Init
func (g *Gom) Init(config MongoConfig) {
	g.Mongo.SetConfig(config)
	g.Mongo.SetContextTimeout(30)
	g.Mongo.SetClient()
}

// Set = Get set query with gom
func (g *Gom) Set() *Set {
	s := NewSet(g)

	return s
}

// ObjectIDFromHex = make object id from hex
func (g *Gom) ObjectIDFromHex(s string) primitive.ObjectID {
	var oid [12]byte

	o, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		return oid
	}

	copy(oid[:], o[:])

	return oid
}

// CheckClient = Check connection successfull or not
func (g *Gom) CheckClient() error {
	err := g.Mongo.Client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		return errors.New(toolkit.Sprintf("Couldn't connect to database : %s", err.Error()))
	}

	toolkit.Println(toolkit.Sprintf("Connected to database: %s", g.Mongo.ConnectionString))

	return nil
}

// GetContext = GetContext for command
func (g *Gom) GetContext() (context.Context, context.CancelFunc) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), g.Mongo.ContextTimeout*time.Second)

	return ctx, cancelFunc
}
