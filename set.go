package gom

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/eaciit/toolkit"
)

// Set = Set struct
type Set struct {
	tableName string
	result    interface{}
	gom       *Gom
	filter    interface{}
	pipe      []bson.M
	sort      struct {
		field  string
		sortBy int
	}
	skip  *int
	limit *int
}

// NewSet = init new set
func NewSet(gom *Gom) *Set {
	s := new(Set)
	s.gom = gom
	s.filter = bson.M{}
	s.pipe = nil
	s.skip = nil
	s.limit = nil

	return s
}

// Table = set table name
func (s *Set) Table(tableName string) *Set {
	s.tableName = tableName

	return s
}

// Result = set target of result
func (s *Set) Result(result interface{}) *Set {
	s.result = result

	return s
}

// Skip = set skip data
func (s *Set) Skip(skip int) *Set {
	s.skip = &skip

	return s
}

// Limit = set limit data
func (s *Set) Limit(limit int) *Set {
	s.limit = &limit

	return s
}

// Sort = set sort data
func (s *Set) Sort(field, sortBy string) *Set {
	s.sort.field = field
	s.sort.sortBy = -1

	if strings.ToLower(sortBy) == "asc" {
		s.sort.sortBy = 1
	}

	return s
}

// Filter = set filter data
func (s *Set) Filter(filter *Filter) *Set {

	if filter != nil {
		main := toolkit.M{}
		inside := toolkit.M{}

		switch filter.Op {
		case OpAnd, OpOr:
			insideArr := []toolkit.M{}
			for _, e := range filter.Items {
				if string(e.Op) == OpSort {
					toolkit.Println("Hmmm", filter.Op)
					insideArr = append(insideArr, toolkit.M{
						string(e.Op): toolkit.M{
							e.Field: e.Value,
						},
					})
				} else {
					insideArr = append(insideArr, toolkit.M{
						e.Field: toolkit.M{
							string(e.Op): e.Value,
						},
					})
				}
				inside.Set(string(filter.Op), insideArr)
			}
			main = inside

		case OpEq, OpNe, OpGt, OpGte, OpLt, OpLte, OpIn, OpNin:
			inside.Set(string(filter.Op), filter.Value)
			main.Set(filter.Field, inside)

		case OpSort:
			inside.Set(string(filter.Op), filter.Value)
			main.Set(filter.Field, inside)

		case OpBetween, OpRange:
			gt := 0
			lt := 0

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(int)
				lt = filter.Value.([]interface{})[1].(int)
			}

			main.Set(filter.Field, toolkit.M{
				"$gt": gt,
				"$lt": lt,
			})

		case OpStartWith:
			main.Set(filter.Field, toolkit.M{}.
				Set("$regex", fmt.Sprintf("^%s.*$", filter.Value)).
				Set("$options", "i"))

		case OpEndWith:
			main.Set(filter.Field, toolkit.M{}.
				Set("$regex", fmt.Sprintf("^.*%s$", filter.Value)).
				Set("$options", "i"))

		case OpContains:
			if len(filter.Value.([]string)) > 1 {
				bfs := []interface{}{}
				for _, ff := range filter.Value.([]string) {
					pfm := toolkit.M{}
					pfm.Set(filter.Field, toolkit.M{}.
						Set("$regex", fmt.Sprintf(".*%s.*", ff)).
						Set("$options", "i"))
					bfs = append(bfs, pfm)
				}
				main.Set("$or", bfs)
			} else {
				main.Set(filter.Field, toolkit.M{}.
					Set("$regex", fmt.Sprintf(".*%s.*", filter.Value.([]string)[0])).
					Set("$options", "i"))
			}

		case OpNot:
			// field := filter.Items[0].Field
			// main.Set(field, toolkit.M{}.Set("$not", filter.Items[0].Field))
			// toolkit.Println(toolkit.JsonStringIndent(main, "\n"))

		}

		s.filter = main
	}

	return s
}

// Pipe = set pipe, if this is set => Filter will be ignored
func (s *Set) Pipe(pipe []bson.M) *Set {
	s.pipe = pipe

	return s
}

func (s *Set) buildPipe() []bson.M {
	pipe := []bson.M{}

	if s.pipe != nil {
		pipe = s.pipe
	} else {
		if s.filter != nil {
			pipe = append(pipe, bson.M{
				"$match": s.filter.(toolkit.M),
			})
		}
	}

	if s.sort.field != "" {
		pipe = append(pipe, bson.M{
			"$sort": bson.M{
				s.sort.field: s.sort.sortBy,
			},
		})
	}

	if s.skip != nil {
		pipe = append(pipe, bson.M{
			"$skip": s.skip,
		})
	}

	if s.limit != nil {
		pipe = append(pipe, bson.M{
			"$limit": s.limit,
		})
	}

	return pipe
}

// Get = get data. it'll use Filter as default. if pipe not null => Filter will be ignored
func (s *Set) Get() error {
	tableName := s.tableName
	result := s.result

	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr && resultVal.Kind() != reflect.Slice {
		return errors.New("result argument must be a slice")
	}

	client := s.gom.Mongo.Client

	collection := client.Database(s.gom.Mongo.Config.Database).Collection(tableName)

	var cur *mongo.Cursor
	var err error

	// cur, err = collection.Aggregate(s.gom.Mongo.Context, s.pipe)
	// } else {
	// 	cur, err = collection.Find(s.gom.Mongo.Context, s.filter)
	// }

	cur, err = collection.Aggregate(s.gom.Mongo.Context, s.buildPipe())

	defer cur.Close(s.gom.Mongo.Context)

	if err != nil {
		log.Fatal("Error finding all documents: ", err)
	}

	cur.All(s.gom.Mongo.Context, result)

	return nil
}

// GetOne = get one data. it'll use Filter as default, pipe ignored.
func (s *Set) GetOne() (err error) {
	tableName := s.tableName
	result := s.result

	resultVal := reflect.ValueOf(result)

	// if resultVal.Kind() == reflect.Slice {
	if resultVal.Kind() != reflect.Ptr && resultVal.Kind() != reflect.Slice {
		err = errors.New("result argument must be a slice")
	} else {
		client := s.gom.Mongo.Client

		collection := client.Database(s.gom.Mongo.Config.Database).Collection(tableName)

		var err error

		err = collection.FindOne(s.gom.Mongo.Context, s.filter).Decode(result)

		if err != nil {
			err = errors.New(toolkit.Sprintf("Error finding document: %s", err.Error()))
		}

	}

	return nil
}

// func (s *Set) Insert(data interface{}) {
// 	client := s.gom.Mongo.Client

// 	collection := client.Database(s.gom.Mongo.Config.Database).Collection(tableName)
// 	cur, err := collection.Find(s.gom.Mongo.Context, s.filter)

// 	defer cur.Close(s.gom.Mongo.Context)

// 	if err != nil {
// 		log.Fatal("Error finding all documents: ", err)
// 	}

// 	cur.All(s.gom.Mongo.Context, result)

// 	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
// 	id := res.InsertedID
// }
