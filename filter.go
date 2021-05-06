package gom

import (
	"strings"
)

// FilterOp is string represent enumeration of supported filter command
type FilterOp string

const (
	// OpAnd is AND
	OpAnd FilterOp = "$and"
	// OpOr is OR
	OpOr = "$or"
	// OpNot is Not
	OpNot = "$not"
	// OpEq is Equal
	OpEq = "$eq"
	// OpNe is Not Equal
	OpNe = "$ne"
	// OpGte is Greater than or Equal
	OpGte = "$gte"
	// OpGt is Greater than
	OpGt = "$gt"
	// OpLt is Less than
	OpLt = "$lt"
	// OpLte is Less than or equal
	OpLte = "$lte"
	// OpRange is range from until
	OpRange = "$range"
	// OpContains is Contains
	OpContains = "$contains"
	// OpStartWith is Start with
	OpStartWith = "$startwith"
	// OpEndWith is End with
	OpEndWith = "$endwith"
	// OpIn is In
	OpIn = "$in"
	// OpNin is Not in
	OpNin = "$nin"
	// OpSort is Sort
	OpSort = "$sort"
	// OpBetween is Between (Custom)
	OpBetween = "between"
	// OpExists is Exists
	OpExists = "$exists"
	// OpBetweenEq is Between Equal
	OpBetweenEq = "betweenEq"
	// OpRangeEq is Range Equal
	OpRangeEq = "rangeEq"
	// ElemMatch is Elem Match operator
	OpElemMatch = "$elemMatch"
)

// Filter holding Items, Field, Operation, and Value
type Filter struct {
	Items []*Filter
	Field string
	Op    FilterOp
	Value interface{}
}

// newFilter create new filter with given parameter
func newFilter(field string, op FilterOp, v interface{}, items []*Filter) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = op
	f.Value = v
	if items != nil {
		f.Items = items
	}
	return f
}

// And create new filter with And operation
func And(items ...*Filter) *Filter {
	return newFilter("", OpAnd, nil, items)
}

// Or create new filter with Or operation
func Or(items ...*Filter) *Filter {
	return newFilter("", OpOr, nil, items)
}

// Sort create new filter with Sort operation
func Sort(field string, sortType string) *Filter {
	sort := -1

	if strings.ToLower(sortType) == "asc" {
		sort = 1
	}

	return newFilter(field, OpSort, sort, nil)
}

// Eq create new filter with Eq operation
func Eq(field string, v interface{}) *Filter {
	return newFilter(field, OpEq, v, nil)
}

// Not create new filter with Not operation
func Not(item *Filter) *Filter {
	return newFilter("", OpNot, nil, []*Filter{item})
}

// Ne create new filter with Ne operation
func Ne(field string, v interface{}) *Filter {
	return newFilter(field, OpNe, v, nil)
}

// Gte create new filter with Gte operation
func Gte(field string, v interface{}) *Filter {
	return newFilter(field, OpGte, v, nil)
}

// Gt create new filter with Gt operation
func Gt(field string, v interface{}) *Filter {
	return newFilter(field, OpGt, v, nil)
}

// Lt create new filter with Lt operation
func Lt(field string, v interface{}) *Filter {
	return newFilter(field, OpLt, v, nil)
}

// Lte create new filter with Lte operation
func Lte(field string, v interface{}) *Filter {
	return newFilter(field, OpLte, v, nil)
}

// Range create new filter with Range operation
func Range(field string, from, to interface{}) *Filter {
	f := newFilter(field, OpRange, nil, nil)
	f.Value = []interface{}{from, to}
	return f
}

// Between create new filter with Between operation (Custom)
func Between(field string, gt, lt interface{}) *Filter {
	f := newFilter(field, OpBetween, nil, nil)
	f.Value = []interface{}{gt, lt}
	return f
}

// RangeEq create new filter with Range Equal operation
func RangeEq(field string, from, to interface{}) *Filter {
	f := newFilter(field, OpRangeEq, nil, nil)
	f.Value = []interface{}{from, to}
	return f
}

// BetweenEq create new filter with Between Equal operation (Custom)
func BetweenEq(field string, gte, lte interface{}) *Filter {
	f := newFilter(field, OpBetweenEq, nil, nil)
	f.Value = []interface{}{gte, lte}
	return f
}

// In create new filter with In operation
func In(field string, inValues ...interface{}) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpIn
	f.Value = inValues
	return f
}

// Nin create new filter with Nin operation
func Nin(field string, ninValues ...interface{}) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpNin
	f.Value = ninValues
	return f
}

// Contains create new filter with Contains operation
func Contains(field string, values ...string) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpContains
	f.Value = values
	return f
}

// StartWith create new filter with StartWith operation
func StartWith(field string, value string) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpStartWith
	f.Value = value
	return f
}

// EndWith create new filter with EndWith operation
func EndWith(field string, value string) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpEndWith
	f.Value = value
	return f
}

// Exists match the documents that contain the field
func Exists(field string, value bool) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpExists
	f.Value = value
	return f
}

// Exists match the documents that contain the field
func ElemMatch(field string, filter *Filter) *Filter {
	f := new(Filter)
	f.Field = field
	f.Op = OpElemMatch
	f.Value = filter
	return f
}
