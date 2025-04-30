package usecase

import (
	"backend/domain/model"
	"backend/domain/repository"
	"encoding/json"
)

type EnemyUsecase interface {
	GetAll() ([]model.Enemy, error)
	GetByName(name string) (*model.Enemy, error)
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

func (u *enemyUsecase) GetByName(name string) (*model.Enemy, error) {
	// Redis から取得
	jsonData, err := u.repo.GetByNameFromRedis(name)
	if err == nil {
		var enemy model.Enemy
		if err := json.Unmarshal([]byte(jsonData), &enemy); err == nil {
			return &enemy, nil
		}
	}

	// MySQL から取得
	enemy, err := u.repo.GetByName(name)
	if err != nil {
		return nil, err
	}

	// Redis にキャッシュ
	if jsonBytes, err := json.Marshal(enemy); err == nil {
		_ = u.repo.SetEnemyToRedis(name, string(jsonBytes))
	}

	return enemy, nil
}
