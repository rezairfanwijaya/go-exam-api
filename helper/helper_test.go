package helper_test

import (
	"errors"
	"testing"

	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	testCases := []struct {
		Name        string
		Path        string
		Expectation interface{}
		WantError   bool
	}{
		{
			Name: "success",
			Path: "../.env",
			Expectation: map[string]string{
				"DATABASE_HOST":     "127.0.0.1",
				"DATABASE_NAME":     "go_exam",
				"DATABASE_PORT":     "3390",
				"DATABASE_USERNAME": "root",
				"DATABASE_PASSWORD": "12345",
			},
			WantError: false,
		}, {
			Name:        "invalid path",
			Path:        "../../.env",
			Expectation: errors.New("missing path, file doesnt exist"),
			WantError:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual, err := helper.GetEnv(testCase.Path)
			if testCase.WantError {
				assert.Equal(t, testCase.Expectation, err)
			} else {
				assert.Equal(t, testCase.Expectation, actual)
			}
		})
	}
}
