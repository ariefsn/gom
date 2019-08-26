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
func (c *Command) Get() (int64, int64, error) {
	tableName := c.set.tableName
	result := c.set.result

	if result == nil {
		return 0, 0, errors.New("result argument must be set")
	}

	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr && resultVal.Kind() != reflect.Slice {
		return 0, 0, errors.New("result argument must be a slice")
	}

	client := c.set.gom.Mongo.Client

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(tableName)

	var cur *mongo.Cursor
	var err error

	cur, err = collection.Aggregate(ctx, c.set.buildPipe())

	defer cur.Close(ctx)

	if err != nil {
		return 0, 0, errors.New(toolkit.Sprintf("Error finding all documents: %s", err.Error()))
	}

	cur.All(ctx, result)

	countTotal, _ := collection.EstimatedDocumentCount(ctx)

	countFilter := int64(0)

	if len(c.set.pipe) == 0 {
		countFilter, _ = collection.CountDocuments(ctx, c.set.filter)
	} else {
		f := bson.M{}
		for _, e := range c.set.buildPipe() {
			if e["$match"] != nil {
				f = e["$match"].(bson.M)
				break
			}
		}
		countFilter, _ = collection.CountDocuments(ctx, f)
	}

	return countFilter, countTotal, nil
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

	client := c.set.gom.Mongo.Client

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(tableName)

	err := collection.FindOne(ctx, c.set.filter).Decode(c.set.result)

	if err != nil {
		return errors.New(toolkit.Sprintf("Error finding document: %s", err.Error()))
	}

	return nil
}

// Insert = insert one data, for multiple data use InsertAll
func (c *Command) Insert(data interface{}) (interface{}, error) {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	dataM, err := c.set.buildData(data, true)

	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	res, err := collection.InsertOne(ctx, dataM)

	if err != nil {
		return nil, err
	}

	id := res.InsertedID

	toolkit.Println("Inserted ID:", id)
	return id, nil
}

// InsertAll = insert multiple data
func (c *Command) InsertAll(data interface{}) ([]interface{}, error) {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	datas, err := c.set.buildData(data, true)

	if err != nil {
		return []interface{}{}, err
	}

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	res, err := collection.InsertMany(ctx, datas.([]interface{}))

	if err != nil {
		return []interface{}{}, err
	}

	ids := res.InsertedIDs

	toolkit.Println("Inserted IDs:", ids)
	return ids, nil
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

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	res, err := collection.UpdateOne(ctx, c.set.filter, bson.M{
		"$set": dataM,
	})

	if err != nil {
		return err
	}

	toolkit.Println("Documents updated:", res.MatchedCount)

	return nil
}

// DeleteOne = delete one data with filter or pipe
func (c *Command) DeleteOne() error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	if len(c.set.filter.(bson.M)) == 0 {
		return errors.New("filter can't be empty")
	}

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	res, err := collection.DeleteOne(ctx, c.set.filter)

	if err != nil {
		return err
	}

	toolkit.Println("Document deleted: ", res.DeletedCount)

	return nil
}

// DeleteAll = delete all data with filter or pipe
func (c *Command) DeleteAll() (int64, error) {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	res, err := collection.DeleteMany(ctx, c.set.filter)

	if err != nil {
		return 0, err
	}

	toolkit.Println("Documents deleted: ", res.DeletedCount)

	return res.DeletedCount, nil
}

// Drop = drop table/collection
func (c *Command) Drop() error {
	client := c.set.gom.Mongo.Client

	collection := client.Database(c.set.gom.Mongo.Config.Database).Collection(c.set.tableName)

	ctx, cancelFunc := c.set.gom.GetContext()
	defer cancelFunc()

	err := collection.Drop(ctx)

	if err != nil {
		return err
	}

	toolkit.Println(toolkit.Sprintf("Collection %s has been deleted", c.set.tableName))

	return nil
}
