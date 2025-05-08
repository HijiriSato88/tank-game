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

// MySQL からランキングを取得する : limit=1万だとindexを貼っても10万行(ALL)読み取ってた
// func (r *rankingRepository) GetRanking(limit int) ([]model.RankingEntry, error) {
// 	var rankings []model.RankingEntry
// 	err := db.DB.Select(&rankings, `
// 		SELECT id, username, high_score
// 		FROM users
// 		WHERE high_score > 0
// 		ORDER BY high_score DESC
// 		LIMIT ?
// 	`, limit)
// 	return rankings, err
// }

func (r *rankingRepository) GetRanking(limit int) ([]model.RankingEntry, error) {
	zs, err := db.Redis.ZRevRangeWithScores(ctx, "ranking", 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	var result []model.RankingEntry
	for i, z := range zs {
		username, ok := z.Member.(string)
		if !ok {
			continue
		}
		result = append(result, model.RankingEntry{
			Username:  username,
			HighScore: int(z.Score),
			Rank:     i + 1,
		})
	}
	
	return result, nil
}

func (r *rankingRepository) ZAddScore(username string, score int) error {
	return db.Redis.ZAdd(ctx, "ranking", redis.Z{
		Score:  float64(score),
		Member: username,
	}).Err()
}
