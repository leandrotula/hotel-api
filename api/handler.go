package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/store"
	"github.com/leandrotula/hotelapi/types"
)

type UserHandler interface {
	HandleGetUser(c *fiber.Ctx) error
	HandleGetUsers(ctx *fiber.Ctx) error
}

func NewUserHandler(store store.UserStore) *UserApiHandler {
	return &UserApiHandler{
		storeUser: store,
	}
}

type UserApiHandler struct {
	storeUser store.UserStore
}

func (u *UserApiHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := u.storeUser.GetUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserApiHandler) HandleGetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(types.User{
		ID:        "1",
		FirstName: "foo",
		LastName:  "bar",
	})

}
