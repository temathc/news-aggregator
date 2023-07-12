package models

import "time"

type Publications struct {
	GUID        *string   `json:"guid,omitempty" db:"guid"`
	Title       string    `json:"title"          db:"title"`
	Description string    `json:"description"    db:"description"`
	PubTime     time.Time `json:"pubtime"        db:"pubtime"`
	Link        string    `json:"link"           db:"link"`
}
