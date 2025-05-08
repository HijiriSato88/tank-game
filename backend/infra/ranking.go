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

func (r *rankingRepository) GetEventBySlug(slug string) (*model.Event, error) {
	var event model.Event
	err := db.DB.Get(&event, `
		SELECT id, slug, end_at
		FROM events
		WHERE slug = ?
	`, slug)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *rankingRepository) GetRanking(redisKey string, limit int) ([]model.RankingEntry, error) {
	zs, err := db.Redis.ZRevRangeWithScores(ctx, redisKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	var result []model.RankingEntry
	for i, z := range zs {
		username, ok := z.Member.(string)
		if !ok {
			continue
		}
		adjustedScore := z.Score
		originalScore := int(adjustedScore / 1e9)

		result = append(result, model.RankingEntry{
			Username:  username,
			HighScore: originalScore,
			Rank:      i + 1,
		})
	}

	return result, nil
}

func (r *rankingRepository) ZAddScore(redisKey string, username string, adjustedScore int64) error {
	return db.Redis.ZAdd(ctx, redisKey, redis.Z{
		Score:  float64(adjustedScore),
		Member: username,
	}).Err()
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
