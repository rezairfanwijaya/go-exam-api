package user

import "errors"

// interface
type IUserService interface {
	Login(input UserInputLogin) (User, error)
}

type UserService struct {
	UserRepo IUserRepository
}

// func new service
func NewService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo}
}

// function to call repo for find user by email
func (s *UserService) Login(input UserInputLogin) (User, error) {
	// cek apakah email sudah terdaftar
	user, _ := s.UserRepo.FindByEmail(input.Email)

	// jika tidak ada
	if user.ID == 0 {
		return user, errors.New("email tidak terdaftar")
	}

	// cek apakah password sama
	if input.Password != user.Password {
		return user, errors.New("password salah")
	}

	// sukses login
	return user, nil
}
