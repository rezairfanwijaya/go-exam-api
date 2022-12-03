package answer

type IAnswerService interface {
	Save(inputs Answers) error
}

type AnswerService struct {
	anserRepo IAnswerRepository
}

func NewService(anserRepo IAnswerRepository) *AnswerService {
	return &AnswerService{anserRepo}
}

// implementasi
func (s *AnswerService) Save(inputs Answers) error {
	for _, input := range inputs.Answers {
		var answer Answer
		answer.Answer = input.Answer
		answer.QuestionID = input.QuestionID

		func(answer Answer) {
			s.anserRepo.Save(answer)
		}(answer)
	}

	return nil
}
