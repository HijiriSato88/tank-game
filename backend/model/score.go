package model

import "time"

type Score struct {
	ID         int       `db:"id"`
	UserID     int       `db:"user_id"`
	Score      int       `db:"score"`
	RecordedAt time.Time `db:"recorded_at"`
	CreatedAt  time.Time `db:"created_at"`
}
