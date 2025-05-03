package model

type RankingEntry struct {
	UserID    int    `db:"id" json:"user_id"`
	Username  string `db:"username" json:"username"`
	HighScore int    `db:"high_score" json:"high_score"`
}
