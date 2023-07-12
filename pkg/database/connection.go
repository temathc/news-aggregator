package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnect(conn string) *sqlx.DB {
	base, err := sqlx.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	return base
}

type RepoDB struct {
	db *sqlx.DB
}

func NewRepoDB(db *sqlx.DB) RepoDB {
	return RepoDB{db}
}
