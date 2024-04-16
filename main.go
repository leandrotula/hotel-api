package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/api"
)

const (
	defaultPort = "5000"
	flagVarName = "port"
)

func main() {

	givenPort := flag.String(flagVarName, defaultPort, "Port to serve on")
	flag.Parse()
	appServer := fiber.New()
	appServer.Get("/users", api.HandleGetUsers)

	err := appServer.Listen(*givenPort)
	if err != nil {
		panic(err)
	}
}
