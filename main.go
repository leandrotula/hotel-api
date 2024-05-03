package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/api"
	"github.com/leandrotula/hotelapi/store"
	"github.com/leandrotula/hotelapi/util"
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
		panic(err)
	}
	userHandler := api.NewUserHandler(store.NewMongoUserStore(client, util.NewEncryptionService()))
	givenPort := flag.String(flagVarName, defaultPort, "Port to serve on")
	flag.Parse()
	appServer := fiber.New()
	appServer.Get("/v1/user/:id", userHandler.HandleGetUser)
	appServer.Delete("/v1/user/:id", userHandler.HandleDeleteUser)
	appServer.Get("/v1/user", userHandler.HandleGetUsers)
	appServer.Post("/v1/user", userHandler.HandleCreateUser)
	appServer.Put("/v1/user/:id", userHandler.HandleUpdateUser)

	err = appServer.Listen(*givenPort)
	if err != nil {
		panic(err)
	}
}
