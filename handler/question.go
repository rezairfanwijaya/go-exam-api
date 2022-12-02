package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/middleware"
	"github.com/rezairfanwijaya/go-exam-api.git/question"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type QuestionHandler struct {
	serviceQuestion question.IQuestionService
	authService     middleware.IAuthService
	userService     user.IUserService
}

const (
	GURU = "guru"
)

func NewHandlerQuestion(serviceQuestion question.IQuestionService,
	authService middleware.IAuthService,
	userService user.IUserService) *QuestionHandler {
	return &QuestionHandler{serviceQuestion, authService, userService}
}

// implementasi
func (h *QuestionHandler) CreateQuestion(c echo.Context) error {
	// cek akun
	// yang boleh membuat soal hanya guru
	// ambil user yang sedang login
	currentUser := c.Get("currentUser").(user.User)

	// panggil service untuk mencari user
	userLoggedin, err := h.userService.GetUserByID(currentUser.ID)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal mengambil data user",
			http.StatusInternalServerError,
			err.Error(),
		)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// cek apakah gutu atau bukan
	if userLoggedin.Role != GURU {
		response := helper.ResponseAPI(
			"Unauthorized",
			"error",
			http.StatusUnauthorized,
			"akses ditolak",
		)

		return c.JSON(http.StatusUnauthorized, response)
	}

	var input question.QuestionCreateInput

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
	questionCreated, err := h.serviceQuestion.Save(input)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal menyimpan question",
			http.StatusBadRequest,
			err.Error(),
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	// format question
	questionFormatted := helper.QuestionFormating(questionCreated)
	response := helper.ResponseAPI(
		"sukses",
		"sukses menyimpan question",
		http.StatusOK,
		questionFormatted,
	)

	return c.JSON(http.StatusOK, response)

}

func (h *QuestionHandler) GetQuestionById(c echo.Context) error {
	// ambil id dari param uri
	param := c.Param("id")

	// konversi ke integer
	id, err := strconv.Atoi(param)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal melakukan konversi id",
			http.StatusInternalServerError,
			err.Error(),
		)

		return c.JSON(http.StatusInternalServerError, response)
	}

	// panggil service
	question, err := h.serviceQuestion.GetByID(id)
	if err != nil {
		response := helper.ResponseAPI(
			"gagal",
			"gagal mengambil question",
			http.StatusBadRequest,
			err.Error(),
		)

		return c.JSON(http.StatusBadRequest, response)
	}

	// format question
	questionFormatted := helper.QuestionFormating(question)
	response := helper.ResponseAPI(
		"sukses",
		"sukses mengambil question",
		http.StatusOK,
		questionFormatted,
	)

	return c.JSON(http.StatusInternalServerError, response)
}
