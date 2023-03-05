package web

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) ProjectPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectPageHandler")
	}
	project, err := web.s.GetProjectByID(id)
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectPageHandler")
	}

	tags, err := web.s.GetTagsMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlePageHandler")
	}

	teamMembers, err := web.s.GetTeamMembersMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlePageHandler")
	}

	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlePageHandler")
	}

	return templates.ProjectPage(project, tags, teamMembers, topTags), nil
}
