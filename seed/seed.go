package main

import (
	"context"
	"github.com/leandrotula/hotelapi/store"
	"github.com/leandrotula/hotelapi/types"
	"github.com/leandrotula/hotelapi/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	uriMongoDb = "mongodb://localhost:27017"
)

func main() {
	log.Println("Running seed process")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriMongoDb))
	if err != nil {
		panic(err)
	}

	userStore := store.NewMongoUserStore(client, util.NewEncryptionService())

	userOne := &types.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@doe.com",
		Password:  "password",
	}

	userTwo := &types.User{
		FirstName: "Mc",
		LastName:  "Giver",
		Email:     "mcgiver@giver.com",
		Password:  "password",
	}

	users := []*types.User{userOne, userTwo}

	usersInserted, err := userStore.InsertMany(context.Background(), users)
	if err != nil {
		panic(err)
	}

	log.Printf("Inserted %v users", len(usersInserted))

}
