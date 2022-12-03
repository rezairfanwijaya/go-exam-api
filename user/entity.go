package user

import (
	"github.com/rezairfanwijaya/go-exam-api.git/answer"
)

// model user
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Answer   answer.Answer
}
