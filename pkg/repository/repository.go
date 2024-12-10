package repository

import (
	"github.com/jmoiron/sqlx"
)

type UrlShortener interface {
	// CreateUser(user types.User ) (int, error)
	// GetUser(UserName, Password string) ( error)
}

type Repository struct {
	UrlShortener
}



func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UrlShortener: NewUrlShortener(db),
	}
}
