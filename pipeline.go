package gom

import "go.mongodb.org/mongo-driver/bson"

// PipeSortParams = params model for pipe sort multiple
type PipeSortParams struct {
	Field     string
	Ascending bool
}

// PipeUnwind = create pipe for unwind arrays. To spesify, prefix the field with dollar sign ($)
func PipeUnwind(path string, showEmptyArrays bool) bson.M {
	m := bson.M{
		"$unwind": bson.M{
			"path":                       path,
			"preserveNullAndEmptyArrays": showEmptyArrays,
		},
	}

	return m
}

// PipeMatch = create pipe for match filter.
func PipeMatch(filter *Filter) bson.M {
	m := bson.M{
		"$match": BuildFilter(filter),
	}

	return m
}

// PipeLookup = create pipe for lookup to another collection.
func PipeLookup(fromCollection, localField, foreignField, as string) bson.M {
	m := bson.M{
		"$lookup": bson.M{
			"from":         fromCollection,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as,
		},
	}

	return m
}

// PipeLimit = create pipe for limit aggregation.
func PipeLimit(limit int) bson.M {
	m := bson.M{
		"$limit": limit,
	}

	return m
}

// PipeSkip = create pipe for skip aggregation.
func PipeSkip(skip int) bson.M {
	m := bson.M{
		"$skip": skip,
	}

	return m
}

// PipeSort = create pipe for single sort aggregation.
func PipeSort(field string, asc bool) bson.M {
	s := 1

	if !asc {
		s = -1
	}

	m := bson.M{
		"$sort": bson.M{
			field: s,
		},
	}

	return m
}

// PipeSortMultiple = create pipe for multiple sort aggregation.
func PipeSortMultiple(sortParams ...PipeSortParams) bson.M {
	sortM := bson.M{}

	for _, p := range sortParams {
		s := 1

		if !p.Ascending {
			s = -1
		}

		sortM[p.Field] = s
	}

	m := bson.M{
		"$sort": sortM,
	}

	return m
}

// PipeProject = create pipe for project aggregation.
func PipeProject(project bson.M) bson.M {
	m := bson.M{
		"$project": project,
	}

	return m
}
