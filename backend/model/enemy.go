package model

import "time"

type Enemy struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	HP          int       `db:"hp" json:"hp"`
	MoveSpeed   float64   `db:"move_speed" json:"move_speed"`
	Score		int		  `db:"score" json:"score"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
