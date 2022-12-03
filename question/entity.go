package question

import "github.com/rezairfanwijaya/go-exam-api.git/answer"

type Question struct {
	ID       int           `json:"id" gorm:"primaryKey"`
	Question string        `json:"question"`
	Answer   answer.Answer `json:"answer"`
}
