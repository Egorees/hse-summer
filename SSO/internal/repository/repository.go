package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func BuildRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
