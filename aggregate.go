package gom

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// BuildAggregate = Build aggregate from filter
func BuildAggregate(filter *Filter) bson.M {
	main := bson.M{}
	inside := []interface{}{}

	switch filter.Op {
	case OpAnd, OpOr:
		insideArr := []interface{}{}

		for _, fi := range filter.Items {
			fRes := BuildAggregate(fi)
			insideArr = append(insideArr, fRes)
		}

		main[string(filter.Op)] = insideArr

	// nin, sort, exists not available in aggregation
	case OpEq, OpNe, OpGt, OpGte, OpLt, OpLte, OpIn, OpNin, OpSort, OpExists:
		inside = append(inside, filter.Field, filter.Value)
		main[string(filter.Op)] = inside

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

			main["$and"] = []bson.M{
				{
					"$gt": []interface{}{filter.Field, gt},
				},
				{
					"$lt": []interface{}{filter.Field, lt},
				},
			}

		case time.Time:
			gt := time.Now()
			lt := time.Now()

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(time.Time)
				lt = filter.Value.([]interface{})[1].(time.Time)
			}

			main["$and"] = []bson.M{
				{
					"$gt": []interface{}{filter.Field, gt},
				},
				{
					"$lt": []interface{}{filter.Field, lt},
				},
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

			main["$and"] = []bson.M{
				{
					"$gte": []interface{}{filter.Field, gt},
				},
				{
					"$lte": []interface{}{filter.Field, lt},
				},
			}

		case time.Time:
			gt := time.Now()
			lt := time.Now()

			if filter.Value != nil {
				gt = filter.Value.([]interface{})[0].(time.Time)
				lt = filter.Value.([]interface{})[1].(time.Time)
			}

			main["$and"] = []bson.M{
				{
					"$gte": []interface{}{filter.Field, gt},
				},
				{
					"$lte": []interface{}{filter.Field, lt},
				},
			}
		}

	case OpStartWith:
		main["$regexMatch"] = bson.M{
			"input":    filter.Field,
			"$regex":   fmt.Sprintf("^%s.*$", filter.Value),
			"$options": "i",
		}

	case OpEndWith:
		main["$regexMatch"] = bson.M{
			"input":    filter.Field,
			"$regex":   fmt.Sprintf("^.*%s$", filter.Value),
			"$options": "i",
		}

	case OpContains:
		if len(filter.Value.([]string)) > 1 {
			bfs := []interface{}{}
			for _, ff := range filter.Value.([]string) {
				pfm := bson.M{}
				pfm["$regexMatch"] = bson.M{
					"input":    filter.Field,
					"$regex":   fmt.Sprintf(".*%s.*", ff),
					"$options": "i",
				}

				bfs = append(bfs, pfm)
			}
			main["$or"] = bfs
		} else {
			main["$regexMatch"] = bson.M{
				"input":    filter.Field,
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
