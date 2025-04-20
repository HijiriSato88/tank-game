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