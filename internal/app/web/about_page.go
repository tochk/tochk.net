package web

import (
	"github.com/pkg/errors"
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) AboutPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	topTags, err := web.s.GetTopTags()
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectsPageHandler")
	}
	return templates.AboutPage(topTags), nil
}
