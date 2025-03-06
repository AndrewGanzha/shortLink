package user

import "learnProject/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (repo *UserRepository) CreateUser(user *User) (*User, error) {
	result := repo.Database.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.Where(&user, "email = ?", email).First(&User{})
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
