package repository

import (
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (r *Repository) GetAllTeamMembers() (res []datastruct.TeamMembers, err error) {
	err = r.db.Select(&res, "SELECT id, name FROM team_members")
	return res, errors.Wrap(err, "repository.GetAllTeamMembers")
}
