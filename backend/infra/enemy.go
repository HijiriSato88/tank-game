package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
)

type enemyRepository struct{}

func NewEnemyRepository() repository.EnemyRepository {
	return &enemyRepository{}
}

func (r *enemyRepository) GetAll() ([]model.Enemy, error) {
	enemies := []model.Enemy{}
	err := db.DB.Select(&enemies, `
		SELECT id, name, hp, move_speed, score, created_at, updated_at
		FROM enemies
	`)
	return enemies, err
}

func (r *enemyRepository) GetByName(name string) (*model.Enemy, error) {
	var enemy model.Enemy;
	err := db.DB.Get(&enemy, `
		SELECT id, name, hp, move_speed, score
		FROM enemies
		WHERE name = ?
	`, name)
	return &enemy, err
}
