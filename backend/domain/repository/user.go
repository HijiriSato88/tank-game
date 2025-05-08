package repository

import "backend/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByUsername(username string) (*model.User, error)
	GetByID(id int) (*model.User, error)
	GetUserScore(userID, eventID int) (*model.UserScore, error)
	CreateUserScore(userID, eventID, score int) error
	UpdateUserScore(userID, eventID, score int) error
}
