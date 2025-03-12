package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"learnProject/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.GetByEmail(email)

	if existedUser == nil {
		return "", errors.New("Wrong Credentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	return existedUser.Email, nil
}

func (service *AuthService) Register(email, password, username string) (string, error) {
	existedUer, _ := service.UserRepository.GetByEmail(email)

	if existedUer != nil {
		return "", errors.New("User already exists")
	}

	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if error != nil {
		return "", error
	}

	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     username,
	}

	_, err := service.UserRepository.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
