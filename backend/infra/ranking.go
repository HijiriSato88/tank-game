package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
)

type rankingRepository struct{}

func NewRankingRepository() repository.RankingRepository {
	return &rankingRepository{}
}

func (r *rankingRepository) GetRanking(limit int) ([]model.RankingEntry, error) {
	var rankings []model.RankingEntry
	err := db.DB.Select(&rankings, `
		SELECT id, username, high_score
		FROM users
		WHERE high_score > 0
		ORDER BY high_score DESC
		LIMIT ?
	`, limit)
	return rankings, err
}
