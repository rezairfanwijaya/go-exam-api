package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/answer"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	m "github.com/rezairfanwijaya/go-exam-api.git/middleware"
)

type AnswerHandler struct {
	answerService answer.IAnswerService
	authService   m.IAuthService
}

func NewHanlderAnswer(answerService answer.IAnswerService, authService m.IAuthService) *AnswerHandler {
	return &AnswerHandler{answerService, authService}
}

// hanlder untuk menyimpan soal
func (h *AnswerHandler) Save(c echo.Context) error {
	var input answer.Answers

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

	// panggil service
	h.answerService.Save(input)

	response := helper.ResponseAPI(
		"sukses",
		"sukses menyimpan jawaban",
		http.StatusOK,
		"sukses",
	)

	return c.JSON(http.StatusOK, response)
}
