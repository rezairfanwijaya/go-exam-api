package question

type QuestionCreateInput struct {
	Question string `json:"question" validate:"required"`
}
