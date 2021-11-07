package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) ArticlesPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
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
	articles, err := web.s.GetArticlesByLanguage(LanguageRu)
	if err != nil {
		return "", errors.Wrap(err, "web.ArticlesPageHandler")
	}
	return templates.ArticlesPage(articles, tags, teamMembers, topTags), nil
}
