package helper

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/question"
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

type userFormatingLogin struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

type userFormatingBytoken struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type questionFormat struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
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

// user formating when success login
func UserFormatingLogin(user user.User, tokenJWT string) userFormatingLogin {
	return userFormatingLogin{
		ID:    user.ID,
		Name:  user.FullName,
		Email: user.Email,
		Role:  user.Role,
		Token: tokenJWT,
	}
}

// user formating by jwt
func UserFormatingByJWT(user user.User) userFormatingBytoken {
	return userFormatingBytoken{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

// format response question
func QuestionFormating(question question.Question) questionFormat {
	return questionFormat{
		ID:       question.ID,
		Question: question.Question,
	}
}

// format response questions
func QuestionsFormating(questions []question.Question) []questionFormat {
	var questionsFormated []questionFormat

	for _, question := range questions {
		questionFormated := QuestionFormating(question)
		questionsFormated = append(questionsFormated, questionFormated)
	}

	return questionsFormated
}

func AuthRole(c echo.Context) (role string, id int) {
	currentUser := c.Get("currentUser").(user.User)
	return currentUser.Role, currentUser.ID
}
