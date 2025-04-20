package model

type Score struct {
	ID       int    `db:"id"`
	UserID   int    `db:"user_id"`
	Score    int    `db:"score"`
	PlayedAt string `db:"played_at"`
}
