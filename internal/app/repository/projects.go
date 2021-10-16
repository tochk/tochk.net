package repository

import "tochkru-golang/internal/app/datastruct"

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
	return res, err
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
	return res, err
}
