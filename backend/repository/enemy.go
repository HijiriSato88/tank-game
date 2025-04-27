package repository

import (
	"backend/db"
	"backend/model"
)

func GetAllEnemies() ([]model.Enemy, error) {
	enemies := []model.Enemy{}
	err := db.DB.Select(&enemies, `
		SELECT id, hp, name, move_speed, score
		FROM enemies
	`)
	return enemies, err
}

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
