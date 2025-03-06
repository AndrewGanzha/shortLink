package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"index;unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewUser(email string, name string, password string) *User {
	user := &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	return user
}
