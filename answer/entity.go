package answer

type Answer struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Answer     string `json:"answer"`
	QuestionID int    `json:"question_id"`
	UserID     int    `json:"user_id"`
}
