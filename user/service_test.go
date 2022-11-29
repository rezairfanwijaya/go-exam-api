package user_test

import (
	"errors"
	"testing"

	mocks "github.com/rezairfanwijaya/go-exam-api.git/mocks/user"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

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
			Name:        "failed",
			Email:       "",
			Expectation: user.User{},
			WantError:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			service := mocks.UserService{}

			if testCase.WantError {
				service.On("GetUserByEmail", testCase.Email).Return(testCase.Expectation, errors.New("user not found"))
			} else {
				service.On("GetUserByEmail", testCase.Email).Return(testCase.Expectation, nil)
			}

			actual, err := service.GetUserByEmail(testCase.Email)

			if (err != nil) != testCase.WantError {
				t.Errorf("Service.GetUserByEmail() error = %v, wantErr %v", err, testCase.WantError)
				return
			}

			if actual != testCase.Expectation {
				t.Errorf("Repository.FindByEmail() = %v, want %v", actual, testCase.Expectation)
			}
		})
	}
}
