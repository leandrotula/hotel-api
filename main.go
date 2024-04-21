package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leandrotula/hotelapi/api"
	"github.com/leandrotula/hotelapi/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultPort = ":5000"
	flagVarName = "port"
	uriMongoDb  = "mongodb://localhost:27017"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uriMongoDb))
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to MongoDB: %v", err))
		panic(err)
	}
	userHandler := api.NewUserHandler(store.NewMongoUserStore(client))
	givenPort := flag.String(flagVarName, defaultPort, "Port to serve on")
	flag.Parse()
	appServer := fiber.New()
	appServer.Get("/v1/user/:id", userHandler.HandleGetUser)

	err = appServer.Listen(*givenPort)
	if err != nil {
		log.Fatal(fmt.Errorf("error listening on port %v: %v", *givenPort, err))
		panic(err)
	}
}
