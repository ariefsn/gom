package gom

import "go.mongodb.org/mongo-driver/bson/primitive"

// ObjectIDFromHex = make object id from hex
func ObjectIDFromHex(s string) primitive.ObjectID {
	var oid [12]byte

	o, err := primitive.ObjectIDFromHex(s)

	if err != nil {
		return oid
	}

	copy(oid[:], o[:])

	return oid
}
