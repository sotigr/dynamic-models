package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProductHandle[T any](db *MongoDatabase) *BaseMongoModel[T] {
	return &BaseMongoModel[T]{
		Collection: "products",
		Db:         db,
	}
}

type Product struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string
	Count int64
}
