package gom

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// SetParams = parameters that optionally pass to Set method
type SetParams struct {
	TableName string
	Result    interface{}
	Filter    *Filter
	Pipe      []bson.M
	SortField string
	SortBy    string
	Skip      int
	Limit     int
	Timeout   time.Duration
}
