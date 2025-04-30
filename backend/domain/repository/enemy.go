package repository

import "backend/domain/model"

type EnemyRepository interface {
	GetAll() ([]model.Enemy, error)
}