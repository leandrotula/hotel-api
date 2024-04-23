package store

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leandrotula/hotelapi/types"
	"github.com/leandrotula/hotelapi/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbname         = "hotel"
	userCollection = "users"
)

type MongoUserStore struct {
	mongoClient       *mongo.Client
	collection        *mongo.Collection
	passwordEncryptor util.UserPasswordEncryptor
}

func NewMongoUserStore(mongoClient *mongo.Client, passwordEncryptor util.UserPasswordEncryptor) *MongoUserStore {
	collection := mongoClient.Database(dbname).Collection(userCollection)
	return &MongoUserStore{
		mongoClient:       mongoClient,
		collection:        collection,
		passwordEncryptor: passwordEncryptor,
	}
}

type UserStore interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
	GetAllUsers(ctx context.Context) ([]*types.User, error)
	InsertUser(ctx context.Context, user *types.User) (*types.User, error)
}

func (store *MongoUserStore) GetUser(ctx context.Context, id string) (*types.User, error) {

	var foundUser types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := store.collection.FindOne(ctx, bson.M{"id": oid}).Decode(&foundUser); err != nil {
		return nil, err
	}
	return &foundUser, nil
}

func (store *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {

	newPassword, err := store.passwordEncryptor.Encrypt(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = newPassword
	result, err := store.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	log.Info(fmt.Sprintf("Inserted a new user with id: %s", result.InsertedID))
	return user, nil
}

func (store *MongoUserStore) GetAllUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User
	cur, err := store.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
