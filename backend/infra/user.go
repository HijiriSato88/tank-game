package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *model.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, password_hash)
		VALUES (?, ?)`,
		user.Username, user.PasswordHash,
	)
	return err
}

func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
