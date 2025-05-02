package repository

import "backend/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetByID(id int) (*model.User, error)
	UpdateHighScore(userID int, highScore int) error
}
