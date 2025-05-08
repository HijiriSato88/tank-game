package model

type RankingEntry struct {
	Username   string `db:"username" json:"username"`
	HighScore  int    `db:"high_score" json:"high_score"`
	Rank       int    `json:"rank"`
}
