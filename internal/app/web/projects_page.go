package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) ProjectsPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	tags, err := web.s.GetTagsMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}
	teamMembers, err := web.s.GetTeamMembersMap()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}
	projects, err := web.s.GetProjectsByLanguage(LanguageRu)
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}
	return templates.ProjectsPage(projects, tags, teamMembers, topTags), nil
}
