package model

import "time"

type Event struct {
	ID    int       `db:"id"`
	Slug  string    `db:"slug"`
	EndAt time.Time `db:"end_at"`
}
