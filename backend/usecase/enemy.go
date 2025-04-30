package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
)

type EnemyUsecase interface {
	GetAll() ([]model.Enemy, error)
}

type enemyUsecase struct {
	repo repository.EnemyRepository
}

func NewEnemyUsecase(r repository.EnemyRepository) EnemyUsecase {
	return &enemyUsecase{repo: r}
}

func (u *enemyUsecase) GetAll() ([]model.Enemy, error) {
	return u.repo.GetAll()
}
