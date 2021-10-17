package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) IndexPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	tags, err := web.s.GetTagsMap()
	if err != nil {
		return "",  errors.Wrap(err, "web.IndexPageHandler")
	}
	teamMembers, err := web.s.GetTeamMembersMap()
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPageHandler")
	}
	projects, err := web.s.GetProjectsByLanguage(LanguageRu)
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPageHandler")
	}
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPageHandler")
	}
	return templates.IndexPage(projects, tags, teamMembers, topTags), nil
}
