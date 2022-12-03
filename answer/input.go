package answer

type AnswerInput struct {
	QuestionID int    `json:"question_id" validate:"required"`
	Answer     string `json:"answer" validate:"required"`
}

type Answers struct {
	Answers []AnswerInput `json:"answers" validate:"required"`
}
