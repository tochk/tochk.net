package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Ready() error {
	return errors.Wrap(r.db.Ping(), "repository.Ready")
}
