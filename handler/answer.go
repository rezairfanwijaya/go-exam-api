package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/answer"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	m "github.com/rezairfanwijaya/go-exam-api.git/middleware"
	"github.com/rezairfanwijaya/go-exam-api.git/question"
)

type AnswerHandler struct {
	answerService   answer.IAnswerService
	authService     m.IAuthService
	questionService question.IQuestionService
}

func NewHanlderAnswer(
	answerService answer.IAnswerService,
	authService m.IAuthService,
	questionService question.IQuestionService,
) *AnswerHandler {
	return &AnswerHandler{
		answerService,
		authService,
		questionService}
}

const (
	SISWA = "siswa"
)

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

	// cek apakah id question tesedia
	for _, questionDetail := range input.Answers {
		question, _ := h.questionService.GetByID(questionDetail.QuestionID)
		if question.ID == 0 {
			errMsg := fmt.Sprintf("question dengan id %v tidak ditemukan", questionDetail.QuestionID)
			response := helper.ResponseAPI(
				"gagal",
				"gagal submit jawaban",
				http.StatusBadRequest,
				errMsg,
			)

			return c.JSON(http.StatusBadRequest, response)
		}
	}

	// yang bisa menjawab ujian hanya siswa saja
	role, id := helper.AuthRole(c)
	if role != SISWA {
		response := helper.ResponseAPI(
			"Unauthorized",
			"error",
			http.StatusUnauthorized,
			"akses ditolak, hanya siswa yang dapat mengisi jawaban",
		)

		return c.JSON(http.StatusUnauthorized, response)
	}

	// panggil service
	err := h.answerService.Save(input, id)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal submit jawaban",
			http.StatusBadRequest,
			err.Error(),
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI(
		"sukses",
		"sukses menyimpan jawaban",
		http.StatusOK,
		"sukses",
	)

	return c.JSON(http.StatusOK, response)
}

func (h *AnswerHandler) GetAnswerByUserID(c echo.Context) error {
	// yang bisa menjawab ujian hanya siswa saja
	role, id := helper.AuthRole(c)
	if role != SISWA {
		response := helper.ResponseAPI(
			"Unauthorized",
			"error",
			http.StatusUnauthorized,
			"akses ditolak, hanya siswa yang diperbolehkan",
		)

		return c.JSON(http.StatusUnauthorized, response)
	}

	answers, err := h.answerService.GetByUserID(id)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal submit jawaban",
			http.StatusBadRequest,
			err.Error(),
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ResponseAPI(
		"sukses",
		"sukses mengambil jawaban",
		http.StatusOK,
		answers,
	)

	return c.JSON(http.StatusOK, response)
}

func (h *AnswerHandler) GetAllAnswer(c echo.Context) error {
	// hanya guru yang bisa melihat semua jawaban
	role, _ := helper.AuthRole(c)

	if role != GURU {
		response := helper.ResponseAPI(
			"Unauthorized",
			"error",
			http.StatusUnauthorized,
			"akses ditolak, hanya guru yang diperbolehkan",
		)

		return c.JSON(http.StatusUnauthorized, response)
	}

	answers := h.answerService.GetAllAnswers()
	response := helper.ResponseAPI(
		"sukses",
		"sukses mengambil semua jawaban",
		http.StatusOK,
		answers,
	)

	return c.JSON(http.StatusOK, response)
}
