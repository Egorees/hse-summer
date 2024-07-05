package repository

import "github.com/jmoiron/sqlx"

func BuildRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sqlx.DB
}
