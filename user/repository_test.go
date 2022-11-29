package user_test

import (
	"errors"
	"testing"

	mocks "github.com/rezairfanwijaya/go-exam-api.git/mocks/user"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

func TestFindByEmail(t *testing.T) {
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
			repo := &mocks.IUserRepository{}

			if testCase.WantError {
				repo.On("FindByEmail", testCase.Email).Return(testCase.Expectation, errors.New("user not found"))
			} else {
				repo.On("FindByEmail", testCase.Email).Return(testCase.Expectation, nil)
			}

			actual, err := repo.FindByEmail(testCase.Email)

			if (err != nil) != testCase.WantError {
				t.Errorf("Repository.FindByEmail() error = %v, wantErr %v", err, testCase.WantError)
				return
			}

			if actual != testCase.Expectation {
				t.Errorf("Repository.FindByEmail() = %v, want %v", actual, testCase.Expectation)
			}
		})
	}
}
