package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type UserHandler struct {
	UserService user.IUserService
}

// func new handler
func NewHandler(userService user.IUserService) *UserHandler {
	return &UserHandler{userService}
}

// func handle user login
func (h *UserHandler) Login(c echo.Context) error {
	// deklarasi input
	var input user.UserInputLogin

	// binding input for validate
	if err := c.Bind(&input); err != nil {
		response := helper.ResponseAPI(
			"gagal binding",
			"gagal melakukan binding",
			http.StatusInternalServerError,
			err,
		)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// validation input user
	if err := c.Validate(&input); err != nil {
		myErr := helper.FormatErrorValidate(err)
		response := helper.ResponseAPI(
			"gagal validasi",
			"gagal melakukan validasi",
			http.StatusBadRequest,
			myErr,
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	// service user
	user, err := h.UserService.Login(input)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal melakukan login",
			http.StatusBadRequest,
			err.Error(),
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	userFormatted := helper.UserFormating(user)
	response := helper.ResponseAPI(
		"sukses",
		"sukses login",
		http.StatusOK,
		userFormatted,
	)

	return c.JSON(http.StatusOK, response)
}
