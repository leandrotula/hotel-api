package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/types"
)

func HandleGetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(types.User{
		ID:        "1",
		FirstName: "foo",
		LastName:  "bar",
	})

}
