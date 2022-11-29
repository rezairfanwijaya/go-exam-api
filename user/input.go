package user

type UserInputLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,len=8"`
}
