package user

// model user
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
