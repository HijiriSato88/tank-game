package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type RankingUsecase interface {
	GetRanking(eventSlug string, limit int) ([]model.RankingEntry, error)
}

type rankingUsecase struct {
	rankingRepo repository.RankingRepository
}

func NewRankingUsecase(r repository.RankingRepository) RankingUsecase {
	return &rankingUsecase{rankingRepo: r}
}

func (u *rankingUsecase) GetRanking(eventSlug string, limit int) ([]model.RankingEntry, error) {
	redisKey := "ranking:" + eventSlug
	return u.rankingRepo.GetRanking(redisKey, limit)
}
