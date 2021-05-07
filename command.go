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

// newCommand = create new command
func newCommand(s *Set) *Command {
	c := new(Command)
	c.set = s

	return c
}

// Pipe = Return Pipe Aggregate
func (c *Command) Pipe() []bson.M {
	return c.set.buildPipe()
}

// Get = get data. it'll use Filter as default. if pipe not null => Filter will be ignored
func (c *Command) Get() (int64, error) {
	tableName := c.set.tableName
	result := c.set.result

	if result == nil {
		return 0, errors.New("result argument must be set")
	}

	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr && resultVal.Kind() != reflect.Slice {
		return 0, errors.New("result argument must be a slice")
	}

	client := c.set.gom.GetClient()

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	if tableName == "" {
		return 0, errors.New("table name not defined")
	}

	collection := client.Database(c.set.gom.GetDatabase()).Collection(tableName)

	var cur *mongo.Cursor
	var err error

	cur, err = collection.Aggregate(ctx, c.set.buildPipe())

	if err != nil {
		return 0, errors.New(toolkit.Sprintf("Error finding all documents: %s", err.Error()))
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, result)

	if err != nil {
		return 0, errors.New(toolkit.Sprintf("Decode error: %s", err.Error()))
	}

	countTotal, _ := collection.EstimatedDocumentCount(ctx)

	return countTotal, nil
}

// GetOne = get one data. it'll use Filter as default, pipe ignored.
func (c *Command) GetOne() error {
	tableName := c.set.tableName
	result := c.set.result

	if result == nil {
		return errors.New("result argument must be set")
	}

	resultVal := reflect.ValueOf(result)

	if resultVal.Kind() != reflect.Ptr {
		return errors.New("result argument must be a pointer")
	}

	if strings.Contains(reflect.TypeOf(result).String(), "[]") {
		return errors.New("result argument must be a pointer, not a slice")
	}

	client := c.set.gom.GetClient()

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(tableName)

	err := collection.FindOne(ctx, c.set.filter).Decode(c.set.result)

	if err != nil {
		return errors.New(toolkit.Sprintf("Error finding document: %s", err.Error()))
	}

	return nil
}

// Insert = insert one data, for multiple data use InsertAll
func (c *Command) Insert(data interface{}) (interface{}, error) {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	dataM, err := c.set.buildData(data, true)

	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	res, err := collection.InsertOne(ctx, dataM)

	if err != nil {
		return nil, err
	}

	id := res.InsertedID

	return id, nil
}

// InsertAll = insert multiple data
func (c *Command) InsertAll(data interface{}) ([]interface{}, error) {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	datas, err := c.set.buildData(data, true)

	if err != nil {
		return []interface{}{}, err
	}

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	res, err := collection.InsertMany(ctx, datas.([]interface{}))

	if err != nil {
		return []interface{}{}, err
	}

	ids := res.InsertedIDs

	return ids, nil
}

// Update = update data with filter or pipe
func (c *Command) Update(data interface{}) (int64, error) {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	dataM, err := c.set.buildData(data, false)

	if err != nil {
		return 0, err
	}

	if len(c.set.filter.(bson.M)) == 0 {
		return 0, errors.New("filter can't be empty")
	}

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	res, err := collection.UpdateOne(ctx, c.set.filter, bson.M{
		"$set": dataM,
	})

	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

// DeleteOne = delete one data with filter or pipe
func (c *Command) DeleteOne() (int64, error) {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	if len(c.set.filter.(bson.M)) == 0 {
		return 0, errors.New("filter can't be empty")
	}

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	res, err := collection.DeleteOne(ctx, c.set.filter)

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

// DeleteAll = delete all data with filter or pipe
func (c *Command) DeleteAll() (int64, error) {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	res, err := collection.DeleteMany(ctx, c.set.filter)

	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

// Drop = drop table/collection
func (c *Command) Drop() error {
	client := c.set.gom.GetClient()

	collection := client.Database(c.set.gom.GetDatabase()).Collection(c.set.tableName)

	ctx, cancelFunc := c.set.GetContext()
	defer cancelFunc()

	err := collection.Drop(ctx)

	if err != nil {
		return err
	}

	return nil
}
