package auth

import (
	"errors"
	"learnProject/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, username string) (string, error) {
	existedUer, _ := service.UserRepository.GetByEmail(email)

	if existedUer != nil {
		return "", errors.New("User already exists")
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     username,
	}

	_, err := service.UserRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
