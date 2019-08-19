package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Hero struct
type Hero struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"Name" bson:"Name"`
	RealName string             `json:"RealName" bson:"RealName"`
	Age      int                `json:"Age" bson:"Age"`
}

// DummyData = Dummy data of hero
func DummyData() []Hero {
	return []Hero{
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Ironman",
			RealName: "Tony Stark",
			Age:      45,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Spiderman",
			RealName: "Peter Parker",
			Age:      20,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Batman",
			RealName: "Bruce Wayne",
			Age:      46,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Green Arrow",
			RealName: "Oliver Queen",
			Age:      34,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Red Arrow",
			RealName: "Thea Queen",
			Age:      27,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Flash",
			RealName: "Barry Allen",
			Age:      28,
		},
		Hero{
			ID:       primitive.NewObjectID(),
			Name:     "Superman",
			RealName: "Kal-El",
			Age:      43,
		},
	}
}
