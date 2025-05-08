package repository

import "backend/domain/model"

type RankingRepository interface {
	GetRanking(limit int) ([]model.RankingEntry, error)
	ZAddScore(username string, highScore int) error
}