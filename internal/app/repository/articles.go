package repository

import (
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"tochkru-golang/internal/app/datastruct"
)

func (r *Repository) GetArticlesByLanguage(language string) (res []datastruct.Articles, err error) {
	err = r.db.Select(&res, `SELECT id, 
			   title, 
			   text,
       		   short_text,
			   created_at,
               language,
			   updated_at 
		FROM articles 
		WHERE language = $1 
		ORDER BY created_at DESC`, language)
	return res, errors.Wrap(err, "repository.GetArticlesByLanguage")
}

func (r *Repository) GetArticleByID(id int) (res datastruct.Articles, err error) {
	err = r.db.Get(&res, `SELECT id, 
			   title, 
			   text,
       		   short_text,
			   created_at,
               language,
			   updated_at 
		FROM articles 
		WHERE id = $1`, id)
	return res, errors.Wrap(err, "repository.GetArticleByID")
}

func (r *Repository) GetArticlesTags(projectIDs []int) (res []datastruct.ArticlesTags, err error) {
	err = r.db.Select(&res, "SELECT article_id, tag_id FROM articles_tags WHERE article_id = ANY($1)", pq.Array(projectIDs))
	return res, errors.Wrap(err, "repository.GetArticlesTags")
}

func (r *Repository) GetArticlesTeamMembers(projectIDs []int) (res []datastruct.ArticlesTeamMembers, err error) {
	err = r.db.Select(&res, "SELECT article_id, team_member_id FROM articles_team_members WHERE article_id = ANY($1)", pq.Array(projectIDs))
	return res, errors.Wrap(err, "repository.GetArticlesTeamMembers")
}
