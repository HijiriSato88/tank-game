package repository

import "backend/domain/model"

type RankingRepository interface {
	GetRanking(limit int) ([]model.RankingEntry, error)
}