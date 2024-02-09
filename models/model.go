package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoModel[T any] interface {
	Insert(entry T) primitive.ObjectID
	FindId(id primitive.ObjectID) (*T, error)
}

type BaseMongoModel[T any] struct {
	Collection string
	Db         *MongoDatabase
}

type MongoDatabase struct {
	Client   *mongo.Client
	Database string
}

func NewDatabase(conStr string, database string) (*MongoDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(conStr))
	if err != nil {
		return nil, err
	}

	return &MongoDatabase{Database: database, Client: c}, nil
}

func (m *BaseMongoModel[T]) getDB() *mongo.Database {
	return m.Db.Client.Database(m.Db.Database)
}

func (m *BaseMongoModel[T]) Insert(entry T) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := m.getDB().Collection(m.Collection).InsertOne(ctx, entry)

	return res.InsertedID.(primitive.ObjectID), err
}

func (m *BaseMongoModel[T]) FindId(id primitive.ObjectID) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result T
	col := m.getDB().Collection(m.Collection)
	err := col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&result)

	return &result, err
}
