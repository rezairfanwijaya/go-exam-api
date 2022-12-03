package question

import "gorm.io/gorm"

// interface
type IQuestionRepository interface {
	FindAll() ([]Question, error)
	Save(question Question) (Question, error)
	FindByID(id int) (Question, error)
	Update(question Question) (Question, error)
	DeleteByID(id int) error
}

type QuestionRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db}
}

// implementasi
func (r *QuestionRepository) FindAll() ([]Question, error) {
	var questions []Question

	if err := r.db.Find(&questions).Error; err != nil {
		return questions, err
	}

	return questions, nil
}

func (r *QuestionRepository) FindByID(id int) (Question, error) {
	var question Question

	if err := r.db.Where("id = ?", id).Find(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *QuestionRepository) Save(question Question) (Question, error) {
	if err := r.db.Create(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *QuestionRepository) Update(question Question) (Question, error) {
	if err := r.db.Save(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

func (r *QuestionRepository) DeleteByID(id int) error {
	var question Question

	if err := r.db.Where("id = ?", id).Delete(&question).Error; err != nil {
		return err
	}

	return nil
}
