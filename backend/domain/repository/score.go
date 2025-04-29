package repository

type ScoreRepository interface {
	Insert(userID int, score int) error
}