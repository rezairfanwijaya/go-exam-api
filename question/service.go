package question

import (
	"errors"
	"fmt"
)

// interface
type IQuestionService interface {
	GetAll() []Question
	Save(input QuestionCreateInput) (Question, error)
	UpdateByID(input QuestionCreateInput, id int) (Question, error)
	GetByID(id int) (Question, error)
	DeleteByID(id int) error
}

type QuestionService struct {
	RepoQuestion IQuestionRepository
}

func NewService(RepoQuestion IQuestionRepository) *QuestionService {
	return &QuestionService{RepoQuestion}
}

// implementasi
func (s *QuestionService) GetAll() []Question {
	questions, _ := s.RepoQuestion.FindAll()

	return questions
}

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

func (s *QuestionService) UpdateByID(input QuestionCreateInput, id int) (Question, error) {
	// id harus lebih dari sama dengan 1
	if id <= 0 {
		return Question{}, errors.New("id harus lebih dari sama dengan 1")
	}

	// cari question by id
	question, _ := s.RepoQuestion.FindByID(id)

	if question.ID == 0 {
		errMsg := fmt.Sprintf("question dengan id %v tidak ditemukan", id)
		return Question{}, errors.New(errMsg)
	}

	// binding struct
	questionInput := Question{
		Question: input.Question,
	}

	question.Question = questionInput.Question

	// panggil repo
	questionUpdated, _ := s.RepoQuestion.Update(question)

	return questionUpdated, nil
}

func (s *QuestionService) DeleteByID(id int) error {
	// id harus lebih dari sama dengan 1
	if id <= 0 {
		return errors.New("id harus lebih dari sama dengan 1")
	}

	// cek apakah ada soal dengan id tersebut
	question, _ := s.RepoQuestion.FindByID(id)

	if question.ID == 0 {
		errMsg := fmt.Sprintf("question dengan id %v tidak ditemukan", id)
		return errors.New(errMsg)
	}
	// panggil repo
	s.RepoQuestion.DeleteByID(id)

	return nil
}
