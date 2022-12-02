package question

import (
	"errors"
	"fmt"
)

// interface
type IQuestionService interface {
	Save(input QuestionCreateInput) (Question, error)
	GetByID(id int) (Question, error)
}

type QuestionService struct {
	repoQuestion IQuestionRepository
}

func NewService(repoQuestion IQuestionRepository) *QuestionService {
	return &QuestionService{repoQuestion}
}

// implementasi
func (s *QuestionService) GetByID(id int) (Question, error) {
	// id harus lebih dari sama dengan 1
	if id <= 0 {
		return Question{}, errors.New("id harus lebih dari sama dengan 1")
	}

	// panggil repository
	question, _ := s.repoQuestion.FindByID(id)

	if question.ID == 0 {
		errMsg := fmt.Sprintf("question dengan id %v tidak ditemukan", id)
		return question, errors.New(errMsg)
	}

	return question, nil
}

func (s *QuestionService) Save(input QuestionCreateInput) (Question, error) {
	// konversi struct
	question := Question{
		Question: input.Question,
	}

	// panggil repo
	questionCreatred, _ := s.repoQuestion.Save(question)

	return questionCreatred, nil
}
