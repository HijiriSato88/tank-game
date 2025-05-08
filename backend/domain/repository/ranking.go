package repository

import "backend/domain/model"

type RankingRepository interface {
	ZAddScore(redisKey, username string, adjustedScore int64) error
	GetRanking(redisKey string, limit int) ([]model.RankingEntry, error)
	GetEventBySlug(slug string) (*model.Event, error)
}