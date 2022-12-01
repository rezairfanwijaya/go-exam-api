package user_test

import (
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

func TestLogin(t *testing.T) {
	type inputUser struct {
		Email    string
		Password string
	}

	testCases := []struct {
		Name        string
		input       inputUser
		Expectation user.User
		WantError   bool
	}{
		{
			Name: "success",
			input: inputUser{
				Email:    "siswapertama@gmail.com",
				Password: "12345678",
			},
			Expectation: user.User{
				ID:       1,
				FullName: "reza",
				Email:    "siswapertama@gmail.com",
				Password: "12345678",
				Role:     "siswa",
			},
			WantError: false,
		}, {
			Name: "wrong email",
			input: inputUser{
				Email:    "root@gmail.com",
				Password: "12345678",
			},
			Expectation: user.User{
				ID:       0,
				FullName: "",
				Email:    "",
				Password: "",
				Role:     "duh",
			},
			WantError: true,
		}, {
			Name: "wrong password",
			input: inputUser{
				Email:    "siswapertama@gmail.com",
				Password: "12345670",
			},
			Expectation: user.User{
				ID:       1,
				FullName: "reza",
				Email:    "siswapertama@gmail.com",
				Password: "12345678",
				Role:     "siswa",
			},
			WantError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock := new(MockIUserRepository)
			userService := user.UserService{mock}

			if testCase.WantError {
				mock.Mock.On("FindByEmail", testCase.input.Email).Return(testCase.Expectation, nil)
				actual, _ := userService.Login(user.UserInputLogin(testCase.input))
				assert.Equal(t, testCase.Expectation, actual)
			} else {
				mock.Mock.On("FindByEmail", testCase.input.Email).Return(testCase.Expectation, nil)
				actual, err := userService.Login(user.UserInputLogin(testCase.input))
				assert.Nil(t, err)
				assert.NotNil(t, actual)
			}
		})
	}
}
