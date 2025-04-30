package repository

import "backend/domain/model"

type EnemyRepository interface {
	GetAll() ([]model.Enemy, error)
	GetByName(name string) (*model.Enemy, error)
	GetByNameFromRedis(name string) (string, error)
	SetEnemyToRedis(name string, jsonStr string) error
}