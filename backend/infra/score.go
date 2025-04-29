package infra

import (
	"backend/db"
	"backend/domain/repository"
)

type scoreRepository struct{}

func NewScoreRepository() repository.ScoreRepository {
	return &scoreRepository{}
}

func (r *scoreRepository) Insert(userID int, score int) error {
	_, err := db.DB.Exec(`
		INSERT INTO scores (user_id, score, recorded_at, created_at)
		VALUES (?, ?, NOW(), NOW())
	`, userID, score)
	return err
}
