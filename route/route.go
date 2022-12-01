package route

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/handler"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/middleware"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

// custom validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRoute(e *echo.Echo, db *gorm.DB) {

	// set validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// service auth
	serviceAuth := middleware.NewServiceAuth()

	// user repo
	repoUser := user.NewRepository(db)
	// user service
	serviceUser := user.NewService(repoUser)
	// user handler
	handlerUser := handler.NewHandler(serviceUser, serviceAuth)

	e.POST("/login", handlerUser.Login)
	e.GET("/home", handlerUser.Home, authMiddleware(serviceAuth, serviceUser))
}

func authMiddleware(authService middleware.IAuthService, userService user.IUserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ambil header authorization
			header := c.Request().Header
			authHeader := header.Get("Authorization")

			// value dari header harus mengandung kata bearer
			if !strings.Contains(authHeader, "Bearer") {
				response := helper.ResponseAPI(
					"Unauthorized",
					"error",
					http.StatusUnauthorized,
					"Masukan string Bearer sebelum token",
				)

				return c.JSON(http.StatusUnauthorized, response)
			}

			// split value authorization untuk menghilangkan kata bearer
			tokenJWT := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenJWT = arrayToken[1]
			}

			// validasi token
			token, err := authService.ValidasiJWT(tokenJWT)
			if err != nil {
				response := helper.ResponseAPI(
					"Unauthorized",
					"gagal melakukan authorization",
					http.StatusUnauthorized,
					err.Error(),
				)

				return c.JSON(http.StatusUnauthorized, response)
			}

			// ambil data dalam token
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := helper.ResponseAPI(
					"Unauthorized",
					"error",
					http.StatusUnauthorized,
					"gagal mengambil data dalam token",
				)

				return c.JSON(http.StatusUnauthorized, response)
			}

			// ambil userid
			userID := int(claim["userID"].(float64))

			// ambil user berdasarkan id
			user, err := userService.GetUserByID(userID)
			if err != nil {
				response := helper.ResponseAPI(
					"Unauthorized",
					"error",
					http.StatusUnauthorized,
					"gagal mengambil user by id",
				)

				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("currentUser", user)
			return nil
		}
	}
}
