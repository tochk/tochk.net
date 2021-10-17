package repository

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (r *Repository) GetProjectsByLanguage(language string) (res []datastruct.Projects, err error) {
	err = r.db.Select(&res, `SELECT id, 
			   title, 
			   short_description,
			   description,
			   image_url, 
			   created_at,
               language,
       		   redirect_url,
			   updated_at 
		FROM projects 
		WHERE language = $1 
		ORDER BY created_at DESC`, language)
	return res, errors.Wrap(err, "repository.GetProjectsByLanguage")
}

func (r *Repository) GetProjectByID(id int) (res datastruct.Projects, err error) {
	err = r.db.Get(&res, `SELECT id, 
			   title, 
			   short_description,
			   description,
			   image_url, 
			   created_at,
               language,
       		   redirect_url,
			   updated_at
		FROM projects 
		WHERE id = $1`, id)
	return res, errors.Wrap(err, "repository.GetProjectByID")
}

func (r *Repository) GetProjectsTags(projectIDs []int) (res []datastruct.ProjectsTags, err error) {
	err = r.db.Select(&res, "SELECT project_id, tag_id FROM projects_tags WHERE project_id = ANY($1)", pq.Array(projectIDs))
	return res, errors.Wrap(err, "repository.GetProjectsTags")
}

func (r *Repository) GetProjectsTeamMembers(projectIDs []int) (res []datastruct.ProjectsTeamMembers, err error) {
	err = r.db.Select(&res, "SELECT project_id, team_member_id FROM projects_team_members WHERE project_id = ANY($1)", pq.Array(projectIDs))
	return res, errors.Wrap(err, "repository.GetProjectsTeamMembers")
}
