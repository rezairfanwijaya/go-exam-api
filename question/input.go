package question

type QuestionCreateInput struct {
	Number   int    `json:"number" validate:"required"`
	Question string `json:"question" validate:"required"`
}
