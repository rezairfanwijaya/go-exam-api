package route

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/handler"
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

	// user repo
	repoUser := user.NewRepository(db)
	// user service
	serviceUser := user.NewService(repoUser)
	// user handler
	handlerUser := handler.NewHandler(serviceUser)

	e.POST("/login", handlerUser.Login)
}
