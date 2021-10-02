package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Ready() error {
	return r.db.Ping()
}
