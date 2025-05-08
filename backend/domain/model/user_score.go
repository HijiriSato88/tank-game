package model

type UserScore struct {
	ID        int `db:"id"`
	UserID    int `db:"user_id"`
	EventID   int `db:"event_id"`
	HighScore int `db:"high_score"`
}
