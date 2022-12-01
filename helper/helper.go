package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type responseAPI struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type userFormating struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

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

// func response api
func ResponseAPI(status, message string, code int, data interface{}) responseAPI {
	meta := meta{
		Status:  status,
		Code:    code,
		Message: message,
	}

	return responseAPI{
		Meta: meta,
		Data: data,
	}
}

// user fortmating
func UserFormating(user user.User) userFormating {
	return userFormating{
		ID:    user.ID,
		Name:  user.FullName,
		Email: user.Email,
		Role:  user.Role,
	}
}
