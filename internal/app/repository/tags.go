package repository

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (r *Repository) GetAllTags() (res []datastruct.Tags, err error) {
	err = r.db.Select(&res, "SELECT id, name FROM tags")
	return res, errors.Wrap(err, "repository.GetAllTags")
}

func (r *Repository) GetTopTags() (res []datastruct.Tags, err error) {
	//todo select top tags
	err = r.db.Select(&res, "SELECT id, name FROM tags")
	return res, errors.Wrap(err, "repository.GetTopTags")
}
