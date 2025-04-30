package repository

import (
	"backend/db"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

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
