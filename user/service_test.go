package user_test

import (
	"errors"
	"testing"

	"github.com/rezairfanwijaya/go-exam-api.git/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockIUserRepository struct {
	Mock mock.Mock
}

func (m MockIUserRepository) FindByEmail(email string) (user.User, error) {
	args := m.Mock.Called(email)
	return args.Get(0).(user.User), args.Error(1)
}

func TestNewService(t *testing.T) {
	mock := new(MockIUserRepository)
	userService := user.NewService(mock)
	assert.NotNil(t, userService)
}

func TestGetUserByEmail(t *testing.T) {
	testCases := []struct {
		Name        string
		Email       string
		Expectation user.User
		WantError   bool
	}{
		{
			Name:  "success",
			Email: "rezairfan@gmail.com",
			Expectation: user.User{
				ID:       1,
				FullName: "reza irfan",
				Email:    "rezairfan@gmail.com",
				Password: "12345678",
				Role:     "siswa",
			},
			WantError: false,
		}, {
			Name:  "failed",
			Email: "",
			Expectation: user.User{
				ID:       0,
				FullName: "",
				Email:    "",
				Password: "",
				Role:     "",
			},
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock := new(MockIUserRepository)
			userService := user.UserService{mock}

			if testCase.WantError {
				mock.Mock.On("FindByEmail", testCase.Email).Return(testCase.Expectation, errors.New("user not found"))
				actual, _ := userService.GetUserByEmail(testCase.Email)
				assert.Equal(t, testCase.Expectation, actual)
			} else {
				mock.Mock.On("FindByEmail", testCase.Email).Return(testCase.Expectation, nil)
				actual, err := userService.GetUserByEmail(testCase.Email)
				assert.Nil(t, err)
				assert.NotNil(t, actual)
			}

		})
	}
}
