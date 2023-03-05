package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) TeamPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}

	teamMembers, err := web.s.GetTeamMembers()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}

	return templates.TeamPage(teamMembers, topTags), nil
}
