package gom

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// BuildFilter = Build filter
func (g *Gom) BuildFilter(filter *Filter) bson.M {
	main := bson.M{}
	inside := bson.M{}

	switch filter.Op {
	case OpAnd, OpOr:
		insideArr := []interface{}{}

		for _, fi := range filter.Items {
			fRes := g.BuildFilter(fi)
			insideArr = append(insideArr, fRes)
		}

		main[string(filter.Op)] = insideArr

	case OpEq, OpNe, OpGt, OpGte, OpLt, OpLte, OpIn, OpNin, OpSort, OpExists:
		inside[string(filter.Op)] = filter.Value
		main[filter.Field] = inside

	// case OpSort:
	// 	inside.Set(string(filter.Op), filter.Value)
	// 	main.Set(filter.Field, inside)

	case OpBetween, OpRange:
		switch filter.Value.([]interface{})[0].(type) {
		case int:
			gt := 0
			lt := 0

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(int)
				lt = filter.Value.([]interface{})[1].(int)
			}

			main[filter.Field] = bson.M{
				"$gt": gt,
				"$lt": lt,
			}
		case time.Time:
			gt := time.Now()
			lt := time.Now()

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(time.Time)
				lt = filter.Value.([]interface{})[1].(time.Time)
			}

			main[filter.Field] = bson.M{
				"$gt": gt,
				"$lt": lt,
			}
		}

	case OpBetweenEq, OpRangeEq:
		switch filter.Value.([]interface{})[0].(type) {
		case int:
			gt := 0
			lt := 0

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(int)
				lt = filter.Value.([]interface{})[1].(int)
			}

			main[filter.Field] = bson.M{
				"$gte": gt,
				"$lte": lt,
			}
		case time.Time:
			gt := time.Now()
			lt := time.Now()

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(time.Time)
				lt = filter.Value.([]interface{})[1].(time.Time)
			}

			main[filter.Field] = bson.M{
				"$gte": gt,
				"$lte": lt,
			}
		}

	case OpStartWith:
		main[filter.Field] = bson.M{
			"$regex":   fmt.Sprintf("^%s.*$", filter.Value),
			"$options": "i",
		}

	case OpEndWith:
		main[filter.Field] = bson.M{
			"$regex":   fmt.Sprintf("^.*%s$", filter.Value),
			"$options": "i",
		}

	case OpContains:
		if len(filter.Value.([]string)) > 1 {
			bfs := []interface{}{}
			for _, ff := range filter.Value.([]string) {
				pfm := bson.M{}
				pfm[filter.Field] = bson.M{
					"$regex":   fmt.Sprintf(".*%s.*", ff),
					"$options": "i",
				}

				bfs = append(bfs, pfm)
			}
			main["$or"] = bfs
		} else {
			main[filter.Field] = bson.M{
				"$regex":   fmt.Sprintf(".*%s.*", filter.Value.([]string)[0]),
				"$options": "i",
			}
		}

	case OpNot:
		// field := filter.Items[0].Field
		// main.Set(field, toolkit.M{}.Set("$not", filter.Items[0].Field))
		// toolkit.Println(toolkit.JsonStringIndent(main, "\n"))

	}

	return main
}
