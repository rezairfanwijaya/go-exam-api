package answer

import "gorm.io/gorm"

type IAnswerRepository interface {
	Save(answer Answer) error
	FindByUserID(userID int) ([]Answer, error)
}

type AnswerRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db}
}

// implementasi
func (r *AnswerRepository) Save(answer Answer) error {
	if err := r.db.Create(&answer).Error; err != nil {
		return err
	}

	return nil
}

func (r *AnswerRepository) FindByUserID(userID int) ([]Answer, error) {
	var answers []Answer
	if err := r.db.Where("user_id = ?", userID).Find(&answers).Error; err != nil {
		return answers, err
	}

	return answers, nil
}
