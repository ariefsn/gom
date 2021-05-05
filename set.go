package gom

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/eaciit/toolkit"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
)

// Set = Set struct
type Set struct {
	tableName      string
	result         interface{}
	gom            *Gom
	filter         interface{}
	pipe           []bson.M
	sortField      *string
	sortBy         *int
	skip           *int
	limit          *int
	command        *Command
	contextTimeout time.Duration
}

// newSet = init new set
func newSet(gom *Gom, params *SetParams) *Set {
	s := new(Set)
	if params == nil {
		s.filter = bson.M{}
		s.pipe = nil
		s.skip = nil
		s.limit = nil
		s.result = nil
		s.tableName = ""
		s.sortField = nil
		s.sortBy = nil
		s.contextTimeout = 30
	} else {
		s.filter = bson.M{}
		if params.Filter != nil {
			s.Filter(params.Filter)
		}

		if params.Pipe != nil {
			s.Pipe(params.Pipe)
		}

		if params.Skip != 0 {
			s.Skip(params.Skip)
		}

		if params.Limit != 0 {
			s.Limit(params.Limit)
		}

		if params.Result != nil {
			s.Result(params.Result)
		}

		if params.TableName != "" {
			s.Table(params.TableName)
		}

		if params.SortField != "" {
			s.Sort(params.SortField, params.SortBy)
		}

		if params.Timeout == 0 {
			s.Timeout(30)
		} else {
			s.Timeout(params.Timeout)
		}
	}

	s.gom = gom
	s.command = newCommand(s)

	return s
}

func (s *Set) reset() {
	s.filter = bson.M{}
	s.limit = nil
	s.pipe = nil
	s.result = nil
	s.skip = nil
	s.sortBy = nil
	s.sortField = nil
	s.tableName = ""
}

// Table = set table/collection name
func (s *Set) Table(tableName string) *Set {
	s.tableName = tableName

	return s
}

// Cmd = choose Command
func (s *Set) Cmd() *Command {
	return s.command
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
	s.sortField = &field
	sort := -1

	if strings.ToLower(sortBy) == "asc" {
		sort = 1
	}

	s.sortBy = &sort

	return s
}

// Filter = set filter data
func (s *Set) Filter(filter *Filter) *Set {

	if filter != nil {
		main := bson.M{}
		inside := bson.M{}

		switch filter.Op {
		case OpAnd, OpOr:
			insideArr := []interface{}{}

			for _, fi := range filter.Items {
				fRes := s.Filter(fi)
				insideArr = append(insideArr, fRes.filter)
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

		case OpElemMatch:
			inside[string(filter.Op)] = s.Filter(filter.Value.(*Filter))
			main[filter.Field] = inside
		}

		s.filter = main
	} else {
		s.filter = bson.M{}
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
				"$match": s.filter.(bson.M),
			})
		}
	}

	if s.sortField != nil {
		pipe = append(pipe, bson.M{
			"$sort": bson.M{
				*s.sortField: s.sortBy,
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

func getValidID(key string) string {
	if key == "ID" || key == "_id" || key == "id" {
		return "_id"
	}

	return key
}

func validateJSONRaw(k string, v json.RawMessage, m bson.M) {
	s := string(v)

	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		m[getValidID(k)] = i
		return
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		m[getValidID(k)] = f
		return
	}
	var t time.Time
	err = json.Unmarshal(v, &t)
	if err == nil {
		m[getValidID(k)] = t
		return
	}
	// 26 => includes double quotes
	if len(s) == 26 {
		var oid primitive.ObjectID
		err = json.Unmarshal(v, &oid)
		if err == nil {
			m[getValidID(k)] = oid
			return
		}
	}
	var objMap map[string]json.RawMessage
	err = json.Unmarshal(v, &objMap)
	if err == nil {
		objMapToBsonM := bson.M{}
		for ko, vo := range objMap {
			validateJSONRaw(ko, vo, objMapToBsonM)
		}

		m[getValidID(k)] = objMapToBsonM
		return
	}
	var slice []json.RawMessage
	err = json.Unmarshal(v, &slice)
	if err == nil {
		tempBsonM := bson.M{}
		validSlice := []interface{}{}
		for _, elSlice := range slice {
			validateJSONRaw(toolkit.RandomString(32), elSlice, tempBsonM)
		}
		for _, vo := range tempBsonM {
			validSlice = append(validSlice, vo)
		}

		m[getValidID(k)] = validSlice
		return
	}
	var itf interface{}
	err = json.Unmarshal(v, &itf)
	if err == nil {
		m[getValidID(k)] = itf
		return
	}
	m[getValidID(k)] = v
}

// buildData = buildData from struct/map to bson M
func (s *Set) buildData(data interface{}, includeID bool) (interface{}, error) {
	var result interface{}
	dataM := bson.M{}

	rv := reflect.ValueOf(data)

	if rv.Kind() != reflect.Ptr {
		return nil, errors.New("data argument must be pointer")
	}

	switch rv.Elem().Kind() {
	case reflect.Struct:
		s, _ := json.Marshal(rv.Interface())

		var mRaw map[string]json.RawMessage

		json.Unmarshal(s, &mRaw)

		for k, v := range mRaw {
			if includeID {
				validateJSONRaw(k, v, dataM)
			} else {
				if k != "_id" {
					validateJSONRaw(k, v, dataM)
				}
			}
		}
		result = dataM

	case reflect.Map:
		v := reflect.ValueOf(rv.Elem().Interface())

		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			if includeID {
				dataM[getValidID(key.String())] = value.Interface()
			} else {
				if key.String() != "_id" {
					dataM[getValidID(key.String())] = value.Interface()
				}
			}
		}

		result = dataM

	case reflect.Slice:
		v := reflect.ValueOf(rv.Elem().Interface())

		datas := make([]interface{}, 0)
		for i := 0; i < v.Len(); i++ {
			value := v.Index(i).Interface()
			datas = append(datas, value)
		}

		result = datas

	default:
		return nil, errors.New("data argument must be a struct or map")
	}

	if result == nil {
		return nil, errors.New("data argument can't be empty")
	}

	return result, nil
}

// Timeout = Timeout for command
func (s *Set) Timeout(seconds time.Duration) *Set {
	if &seconds == nil {
		seconds = 30
	}

	s.contextTimeout = seconds

	return s
}

// GetContext = GetContext for command
func (s *Set) GetContext() (context.Context, context.CancelFunc) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), s.contextTimeout*time.Second)

	return ctx, cancelFunc
}
