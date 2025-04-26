package repository

import (
	"backend/db"
)

func InsertScore(userID int, score int) error {
	_, err := db.DB.Exec(`
		INSERT INTO scores (user_id, score, recorded_at, created_at)
		VALUES (?, ?, NOW(), NOW())
	`, userID, score)
	return err
}
