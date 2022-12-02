package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/middleware"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type UserHandler struct {
	UserService user.IUserService
	AuthService middleware.IAuthService
}

// func new handler
func NewHandlerUser(UserService user.IUserService, AuthService middleware.IAuthService) *UserHandler {
	return &UserHandler{UserService, AuthService}
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
			err.Error(),
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

	// generate jwt
	tokenJwt, err := h.AuthService.GenerateJWT(user)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal membuat token",
			http.StatusInternalServerError,
			err.Error(),
		)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userFormatted := helper.UserFormatingLogin(user, tokenJwt)
	response := helper.ResponseAPI(
		"sukses",
		"sukses login",
		http.StatusOK,
		userFormatted,
	)

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserByTokenJWT(c echo.Context) error {
	// ambil user yang sedang login
	currentUser := c.Get("currentUser").(user.User)

	// panggil service untuk mencari user
	userLoggedin, err := h.UserService.GetUserByID(currentUser.ID)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal mengambil data user",
			http.StatusInternalServerError,
			err.Error(),
		)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// format user
	userFormating := helper.UserFormatingByJWT(userLoggedin)
	response := helper.ResponseAPI(
		"sukses",
		"sukses mengambil data user",
		http.StatusOK,
		userFormating,
	)

	return c.JSON(http.StatusOK, response)
}
