package user

import "errors"

// interface
type IUserService interface {
	GetUserByID(id int) (User, error)
	GetUserByEmail(email string) (User, error)
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

// get user by id
func (s *UserService) GetUserByID(id int) (User, error) {
	user, _ := s.UserRepo.FindByID(id)

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	return user, nil
}

// get user by email
func (s *UserService) GetUserByEmail(email string) (User, error) {
	user, _ := s.UserRepo.FindByEmail(email)

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	return user, nil
}
