package gom

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (d *Gom) Init(config MongoConfig) {
	d.Mongo.SetConfig(config)
	d.Mongo.SetContext(30)
	d.Mongo.SetClient()
}

// Set = Get set query with gom
func (d *Gom) Set() *Set {
	s := NewSet(d)

	return s
}

// ObjectIDFromHex = make object id from hex
func (d *Gom) ObjectIDFromHex(s string) primitive.ObjectID {
	var oid [12]byte

	o, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		return oid
	}

	copy(oid[:], o[:])

	return oid
}
