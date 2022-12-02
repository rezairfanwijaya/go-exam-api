package question

type Question struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Question string `json:"question"`
}
