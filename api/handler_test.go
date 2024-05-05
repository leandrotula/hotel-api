package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrotula/hotelapi/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SuiteTest struct {
	suite.Suite
}

type HandlerMock struct {
	mock.Mock
}

type HandlerInputData struct {
	expectedCode int
	mockedUsers  []types.User
	userId       string
	expectedSize int
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

	err := args.Error(1)
	if len(content) == 0 {
		return nil, err
	}
	return []*types.User{
		{
			FirstName: content[0].FirstName,
			LastName:  content[0].LastName,
			Email:     content[0].Email,
			ID:        content[0].ID,
		},
	}, err

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

func (s *SuiteTest) TestUserApiHandler_HandleGetUser() {

	allScenarios := []HandlerInputData{
		{
			expectedCode: http.StatusOK,
			mockedUsers: []types.User{
				{
					FirstName: "test user2",
					LastName:  "test user2",
					Email:     "test2@test.com",
					ID:        "772ecba995e45cd5628b088f",
				},
			},
			userId: "772ecba995e45cd5628b088f",
		},
		{
			expectedCode: http.StatusOK,
			mockedUsers:  []types.User{{}},
			userId:       "772ecba995e45cd5628b088f",
		},
	}

	for _, data := range allScenarios {

		app := fiber.New()

		storeMock := new(HandlerMock)

		storeMock.On("GetUser", mock.Anything, mock.Anything).Return(&data.mockedUsers[0], nil)

		handlerService := NewUserHandler(storeMock)

		app.Get("/:id", handlerService.HandleGetUser)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/%s", data.userId), nil)

		result, _ := app.Test(req)

		if result != nil {
			assert.NotNil(s.T(), result)
			assert.Equal(s.T(), data.expectedCode, result.StatusCode)
		}

	}

}

func (s *SuiteTest) TestUserApiHandler_HandleGetUsers() {

	allScenarios := []HandlerInputData{
		{
			expectedCode: http.StatusOK,
			mockedUsers: []types.User{
				{
					FirstName: "test user2",
					LastName:  "test user2",
					Email:     "test2@test.com",
					ID:        "772ecba995e45cd5628b088f",
				},
			},
			expectedSize: 1,
		},
		{
			expectedCode: http.StatusOK,
			mockedUsers:  []types.User{},
			expectedSize: 0,
		},
	}

	for _, data := range allScenarios {

		app := fiber.New()

		storeMock := new(HandlerMock)

		storeMock.On("GetAllUsers", mock.Anything).Return(data.mockedUsers, nil)

		handlerService := NewUserHandler(storeMock)

		app.Get("/", handlerService.HandleGetUsers)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		result, _ := app.Test(req)

		if result != nil {
			var users []types.User
			content, _ := io.ReadAll(result.Body)

			err := json.Unmarshal(content, &users)
			assert.NoError(s.T(), err)

			assert.NotNil(s.T(), result)
			assert.Equal(s.T(), data.expectedCode, result.StatusCode)
			assert.Equal(s.T(), data.expectedSize, len(users))
		} else {
			assert.Nil(s.T(), result)
		}

	}

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}
