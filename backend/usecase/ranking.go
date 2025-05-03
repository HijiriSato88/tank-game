package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type RankingUsecase interface {
	GetRanking(limit int) ([]model.RankingEntry, error)
}

type rankingUsecase struct {
	rankingRepo repository.RankingRepository
}

func NewRankingUsecase(r repository.RankingRepository) RankingUsecase {
	return &rankingUsecase{rankingRepo: r}
}

func (u *rankingUsecase) GetRanking(limit int) ([]model.RankingEntry, error) {
	return u.rankingRepo.GetRanking(limit)
}
