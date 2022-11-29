package user

import "errors"

// interface
type UserService interface {
	GetUserByEmail(email string) (User, error)
}

type userService struct {
	userRepo IUserRepository
}

// func new service
func NewService(userRepo IUserRepository) *userService {
	return &userService{userRepo}
}

// function to call repo for find user by email
func (s *userService) GetUserByEmail(email string) (User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
