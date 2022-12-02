package question

type Question struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	NumberQuestion int    `json:"number_question"`
	Question       string `json:"question"`
}
