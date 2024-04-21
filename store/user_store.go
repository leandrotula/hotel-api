package store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os/user"
)

const (
	dbname         = "hotel"
	userCollection = "users"
)

type MongoUserStore struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewMongoUserStore(mongoClient *mongo.Client) *MongoUserStore {
	collection := mongoClient.Database(dbname).Collection(userCollection)
	return &MongoUserStore{
		mongoClient: mongoClient,
		collection:  collection,
	}
}

type UserStore interface {
	GetUser(ctx context.Context, id string) (*user.User, error)
}

func (store *MongoUserStore) GetUser(ctx context.Context, id string) (*user.User, error) {

	var foundUser user.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := store.collection.FindOne(ctx, bson.M{"id": oid}).Decode(&foundUser); err != nil {
		return nil, err
	}
	return &foundUser, nil
}
