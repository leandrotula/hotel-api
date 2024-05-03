package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HandlerMock struct {
	mock.Mock
}

func (h *HandlerMock) GetUser(ctx context.Context, id string) (*types.User, error) {

	args := h.Called()
	return args.Get(0).(*types.User), args.Error(1)
}

func (h *HandlerMock) GetAllUsers(ctx context.Context) ([]*types.User, error) {

	args := h.Called()
	content, ok := args.Get(0).([]types.User)
	if !ok {
		panic("invalid argument")
	}
	return []*types.User{
		{
			FirstName: content[0].FirstName,
			LastName:  content[0].LastName,
			Email:     content[0].Email,
			ID:        content[0].ID,
		},
	}, args.Error(1)

}

func (h *HandlerMock) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {

	return nil, nil

}

func (h *HandlerMock) DeleteUser(ctx context.Context, id string) error {

	return nil

}

func (h *HandlerMock) UpdateUser(ctx context.Context, id string, user *types.User) error {

	return nil
}

func TestUserApiHandler_HandleGetUser(t *testing.T) {

	app := fiber.New()
	defaultId := "662ecba995e45cd5628b088f"
	storeMock := new(HandlerMock)

	storeMock.On("GetUser", mock.Anything, mock.Anything).Return(&types.User{
		FirstName: "test user",
		LastName:  "test user",
		Email:     "test@test.com",
		ID:        defaultId,
	}, nil)

	handlerService := NewUserHandler(storeMock)

	app.Get("/:id", handlerService.HandleGetUser)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", defaultId), nil)

	result, _ := app.Test(req)

	assert.NotNil(t, result)
	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestUserApiHandler_HandleGetUsers(t *testing.T) {

	app := fiber.New()
	storeMock := new(HandlerMock)

	mockUsers := []types.User{
		{
			FirstName: "test user2",
			LastName:  "test user2",
			Email:     "test2@test.com",
			ID:        "772ecba995e45cd5628b088f",
		},
	}
	storeMock.On("GetAllUsers", mock.Anything).Return(mockUsers, nil)

	handlerService := NewUserHandler(storeMock)

	app.Get("/", handlerService.HandleGetUsers)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	result, _ := app.Test(req)

	var users []types.User
	content, _ := io.ReadAll(result.Body)

	err := json.Unmarshal(content, &users)
	assert.NoError(t, err)

	fmt.Println(users)
	assert.NotNil(t, result)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, 1, len(users))
}
