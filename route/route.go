package route

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/answer"
	"github.com/rezairfanwijaya/go-exam-api.git/handler"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/middleware"
	"github.com/rezairfanwijaya/go-exam-api.git/question"
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
	handlerUser := handler.NewHandlerUser(serviceUser, serviceAuth)

	// question repo
	repoQuestion := question.NewRepository(db)
	// question service
	serviceQuestion := question.NewService(repoQuestion)
	// question handler
	handlerQuestion := handler.NewHandlerQuestion(serviceQuestion, serviceAuth, serviceUser)

	// answer repo
	repoAnswer := answer.NewRepository(db)
	// answer service
	serviceAnswer := answer.NewService(repoAnswer)
	// answer handler
	handlerAnswer := handler.NewHanlderAnswer(serviceAnswer, serviceAuth, serviceQuestion)

	// endpoint
	e.POST("/login", handlerUser.Login)

	v1 := e.Group("/v1")
	v1.GET("/user/info", handlerUser.GetUserByTokenJWT, authMiddleware(serviceAuth, serviceUser))

	v1.GET("/questions", handlerQuestion.GetAllQuestion, authMiddleware(serviceAuth, serviceUser))
	v1.GET("/question/:id", handlerQuestion.GetQuestionById, authMiddleware(serviceAuth, serviceUser))
	v1.POST("/question", handlerQuestion.CreateQuestion, authMiddleware(serviceAuth, serviceUser))
	v1.PUT("/question/:id", handlerQuestion.UpdateQuestion, authMiddleware(serviceAuth, serviceUser))
	v1.DELETE("/question/:id", handlerQuestion.DeleteQuestion, authMiddleware(serviceAuth, serviceUser))

	v1.POST("/answer", handlerAnswer.Save, authMiddleware(serviceAuth, serviceUser))

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
			return next(c)
		}
	}
}
