package repository

import (
	"backend/db"
	"backend/model"
)

func CreateUser(user *model.User) error {
	_, err := db.DB.Exec(`
		INSERT INTO users (username, password_hash)
		VALUES (?, ?)`,
		user.Username, user.PasswordHash,
	)
	return err
}

func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}