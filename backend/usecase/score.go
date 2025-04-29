package usecase

import (
	"backend/domain/repository"
	"errors"
)

type ScoreUsecase interface {
	InsertScore(userID int, score int) error
}

type scoreUsecase struct {
	scoreRepo repository.ScoreRepository
}

func NewScoreUsecase(r repository.ScoreRepository) ScoreUsecase {
	return &scoreUsecase{scoreRepo: r}
}

func (u *scoreUsecase) InsertScore(userID int, score int) error {
	if score < 0 {
		return errors.New("invalid score: must be non-negative")
	}
	return u.scoreRepo.Insert(userID, score)
}
