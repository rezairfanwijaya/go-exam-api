package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

// function get env value
func GetEnv(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return nil, errors.New("missing path, file doesnt exist")
	}
	return env, nil
}

// function error validate
func FormatErrorValidate(err error) []string {
	var myErr []string

	for _, e := range err.(validator.ValidationErrors) {
		// generate error from validator
		errMessage := fmt.Sprintf("error on filed: %v, condition: %v", e.Field(), e.ActualTag())
		myErr = append(myErr, errMessage)
	}

	return myErr
}
