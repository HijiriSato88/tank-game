package repository

import (
	"backend/db"
	"backend/domain/model"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetEnemyByName(name string) (*model.Enemy, error) {
	var enemy model.Enemy
	err := db.DB.Get(&enemy, `
		SELECT id, name, hp, move_speed, score
		FROM enemies
		WHERE name = ?
	`, name)
	if err != nil {
		return nil, err
	}
	return &enemy, nil
}

func GetEnemyByNameFromRedis(name string) (string, error) {
	key := fmt.Sprintf("enemy:%s", name)
	jsonData, err := db.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("not found")
	} else if err != nil {
		return "", err
	}

	return jsonData, nil
}
