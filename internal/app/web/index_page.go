package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) IndexPage(w http.ResponseWriter, r *http.Request) (string, error) {
	tags, err := web.s.GetTagsMap()
	if err != nil {
		return "",  errors.Wrap(err, "web.IndexPage")
	}
	teamMembers, err := web.s.GetTeamMembersMap()
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPage")
	}
	projects, err := web.s.GetProjectsByLanguage(LanguageRu)
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPage")
	}
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.IndexPage")
	}
	return templates.IndexPage(projects, tags, teamMembers, topTags), nil
}
