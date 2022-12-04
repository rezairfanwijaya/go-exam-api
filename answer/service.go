package answer

import (
	"errors"
	"fmt"
)

type IAnswerService interface {
	Save(inputs Answers, userID int) error
	GetByUserID(userID int) ([]Answer, error)
	GetAllAnswers() []Answer
	DeleteAnswerByQuestionID(questionID, userID int) error
}

type AnswerService struct {
	answerRepo IAnswerRepository
}

func NewService(answerRepo IAnswerRepository) *AnswerService {
	return &AnswerService{answerRepo}
}

// implementasi
func (s *AnswerService) Save(inputs Answers, userID int) error {
	// question deengan id yang sama tidak boleh di jawab dua kali oleh user yang sama
	answersByUserID, _ := s.answerRepo.FindByUserID(userID)
	for _, answerByUserID := range answersByUserID {
		for _, input := range inputs.Answers {
			if answerByUserID.UserID == userID && input.QuestionID == answerByUserID.QuestionID {
				errMsg := fmt.Sprintf("question dengan id %v sudah dijawab", input.QuestionID)
				return errors.New(errMsg)
			}
		}
	}

	for _, input := range inputs.Answers {
		var answer Answer
		answer.Answer = input.Answer
		answer.QuestionID = input.QuestionID
		answer.UserID = userID

		go func(answer Answer) {
			s.answerRepo.Save(answer)
		}(answer)
	}

	return nil
}

func (s *AnswerService) GetByUserID(userID int) ([]Answer, error) {
	answers, _ := s.answerRepo.FindByUserID(userID)
	return answers, nil
}

func (s *AnswerService) GetAllAnswers() []Answer {
	answers, _ := s.answerRepo.FindAll()
	return answers
}

func (s *AnswerService) DeleteAnswerByQuestionID(questionID, userID int) error {
	// question id tidak boleh < 1
	if questionID < 1 {
		return errors.New("question id harus lebih sama dengan 1")
	}

	// cek aoakah jawaban sudah ada
	answer, _ := s.answerRepo.FindByUserIDAndQuestionID(questionID, userID)
	if answer.ID == 0 {
		errMsg := fmt.Sprintf("anda belum menjawab pertanyaan dengan id %v", questionID)
		return errors.New(errMsg)
	}

	// delete
	s.answerRepo.DeleteByQuestionID(questionID, userID)
	return nil
}
