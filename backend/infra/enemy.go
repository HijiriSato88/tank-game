package infra

import (
	"backend/db"
	"backend/domain/model"
	"backend/domain/repository"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type enemyRepository struct{}

func NewEnemyRepository() repository.EnemyRepository {
	return &enemyRepository{}
}

func (r *enemyRepository) GetAll() ([]model.Enemy, error) {
	enemies := []model.Enemy{}
	err := db.DB.Select(&enemies, `
		SELECT id, name, hp, move_speed, score
		FROM enemies
	`)
	return enemies, err
}

func (r *enemyRepository) GetByName(name string) (*model.Enemy, error) {
	var enemy model.Enemy
	err := db.DB.Get(&enemy, `
		SELECT id, name, hp, move_speed, score
		FROM enemies
		WHERE name = ?
	`, name)
	return &enemy, err
}

func (r *enemyRepository) GetByNameFromRedis(name string) (string, error) {
	key := fmt.Sprintf("enemy:%s", name)
	jsonData, err := db.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("not found")
	} else if err != nil {
		return "", err
	}
	return jsonData, nil
}

func (r *enemyRepository) SetEnemyToRedis(name string, jsonStr string) error {
	key := fmt.Sprintf("enemy:%s", name)
	return db.Redis.Set(ctx, key, jsonStr, 0).Err() // TTL = 0（無期限）。設定したければ第3引数変更。
}
