package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type userHandler struct {
	userService user.UserService
}

// func new handler
func NewHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

// func handle user login
func (h *userHandler) Login(c echo.Context) error {
	// deklarasi input
	var input user.UserInputLogin

	// binding input for validate
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	// validation input user
	if err := c.Validate(&input); err != nil {
		myErr := helper.FormatErrorValidate(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": myErr,
		})
	}

	return c.JSON(http.StatusOK, input)
}
