package gom

import (
	"errors"
	"reflect"
	"strings"

	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Command = command struct
type Command struct {
	set *Set
}

// NewCommand = create new command
func NewCommand(s *Set) *Command {
	c := new(Command)
	c.set = s

	return c
}

// Get = get data. it'll use Filter as default. if pipe not null => Filter will be ignored
func (c *Command) Get() error {
	tableName := c.set.tableName
	result := c.set.result

	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr && resultVal.Kind() != reflect.Slice {
		return errors.New("result argument must be a slice")
	}

	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(tableName)

	var cur *mongo.Cursor
	var err error

	cur, err = collection.Aggregate(c.set.gom.Mongo.Context, c.set.buildPipe())

	defer cur.Close(c.set.gom.Mongo.Context)

	if err != nil {
		return errors.New(toolkit.Sprintf("Error finding all documents: %s", err.Error()))
	}

	cur.All(c.set.gom.Mongo.Context, result)

	return nil
}

// GetOne = get one data. it'll use Filter as default, pipe ignored.
func (c *Command) GetOne() error {
	tableName := c.set.tableName
	result := c.set.result

	resultVal := reflect.ValueOf(result)

	// if resultVal.Kind() == reflect.Slice {
	if resultVal.Kind() != reflect.Ptr {
		return errors.New("result argument must be a pointer")
	}

	if strings.Contains(reflect.TypeOf(result).String(), "[]") {
		return errors.New("result argument must be a pointer, not a slice")
	}

	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(tableName)

	err := collection.FindOne(c.set.gom.Mongo.Context, c.set.filter).Decode(c.set.result)

	if err != nil {
		return errors.New(toolkit.Sprintf("Error finding document: %s", err.Error()))
	}

	return nil
}

// Insert = insert one data, for multiple data use InsertAll
func (c *Command) Insert(data interface{}) error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	dataM, err := c.set.buildData(data, true)

	if err != nil {
		return err
	}

	res, err := collection.InsertOne(c.set.gom.Mongo.Context, dataM)

	if err != nil {
		return err
	}

	id := res.InsertedID

	toolkit.Println("Inserted ID:", id)
	return nil
}

// InsertAll = insert multiple data
func (c *Command) InsertAll(data interface{}) error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	datas, err := c.set.buildData(data, true)

	if err != nil {
		return err
	}

	res, err := collection.InsertMany(c.set.gom.Mongo.Context, datas.([]interface{}))

	if err != nil {
		return err
	}

	ids := res.InsertedIDs

	toolkit.Println("Inserted IDs:", ids)
	return nil
}

// Update = update data with filter or pipe
func (c *Command) Update(data interface{}) error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	dataM, err := c.set.buildData(data, false)

	if err != nil {
		return err
	}

	if len(c.set.filter.(bson.M)) == 0 {
		return errors.New("filter can't be empty")
	}

	res, err := collection.UpdateOne(c.set.gom.Mongo.Context, c.set.filter, bson.M{
		"$set": dataM,
	})

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("Documents not found")
	}

	toolkit.Println("Documents updated:", res.MatchedCount)

	return nil
}

// Delete = delete data with filter or pipe
func (c *Command) Delete() error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	if len(c.set.filter.(bson.M)) == 0 {
		return errors.New("filter can't be empty")
	}

	res, err := collection.DeleteOne(c.set.gom.Mongo.Context, c.set.filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("Documents not found")
	}

	toolkit.Println("Document deleted: ", res.DeletedCount)

	return nil
}
