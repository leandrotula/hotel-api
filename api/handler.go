package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/leandrotula/hotelapi/store"
	"github.com/leandrotula/hotelapi/types"
)

type UserHandler interface {
	HandleGetUser(c *fiber.Ctx) error
	HandleGetUsers(ctx *fiber.Ctx) error
	HandleCreateUser(c *fiber.Ctx) error
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

	allUsers, err := u.storeUser.GetAllUsers(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(allUsers)

}

func (u *UserApiHandler) HandleCreateUser(c *fiber.Ctx) error {

	var user types.CreateUserPayload
	if err := c.BodyParser(&user); err != nil {
		log.Fatal("Failed to parse body")
		return err
	}
	userToInsert := types.NewCreateUser(&user)
	insertedUser, err := u.storeUser.InsertUser(c.Context(), userToInsert)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}
