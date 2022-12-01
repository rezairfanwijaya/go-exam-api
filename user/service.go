package user

import "errors"

// interface
type IUserService interface {
	GetUserByEmail(email string) (User, error)
}

type UserService struct {
	UserRepo IUserRepository
}

// func new service
func NewService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo}
}

// function to call repo for find user by email
func (s *UserService) GetUserByEmail(email string) (User, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
