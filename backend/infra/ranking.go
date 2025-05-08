package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
	"github.com/redis/go-redis/v9"
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

func (r *rankingRepository) ZAddScore(username string, score int) error {
	return db.Redis.ZAdd(ctx, "ranking", redis.Z{
		Score:  float64(score),
		Member: username,
	}).Err()
}
