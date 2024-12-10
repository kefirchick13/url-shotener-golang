package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewUrlShortener(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
