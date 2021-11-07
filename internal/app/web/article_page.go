package web

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) ArticlePageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectPageHandler")
	}
	tags, err := web.s.GetTagsMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlesPageHandler")
	}
	teamMembers, err := web.s.GetTeamMembersMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlesPageHandler")
	}
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlesPageHandler")
	}
	article, err := web.s.GetArticleByID(id)
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlesPageHandler")
	}
	return templates.ArticlePage(article, tags, teamMembers, topTags), nil
}
