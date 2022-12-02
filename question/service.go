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
	RepoQuestion IQuestionRepository
}

func NewService(RepoQuestion IQuestionRepository) *QuestionService {
	return &QuestionService{RepoQuestion}
}

// implementasi
func (s *QuestionService) GetByID(id int) (Question, error) {
	// id harus lebih dari sama dengan 1
	if id <= 0 {
		return Question{}, errors.New("id harus lebih dari sama dengan 1")
	}

	// panggil repository
	question, _ := s.RepoQuestion.FindByID(id)

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
	questionCreatred, _ := s.RepoQuestion.Save(question)

	return questionCreatred, nil
}
