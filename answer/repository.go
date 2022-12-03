package answer

import "gorm.io/gorm"

type IAnswerRepository interface {
	Save(answer Answer) error
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
