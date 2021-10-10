package web

import (
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) IndexPage(w http.ResponseWriter, r *http.Request) (string, error) {
	projects, err := web.s.GetProjectsByLanguage(LanguageRu)
	if err != nil {
		return "", err
	}
	return templates.IndexPage(projects), nil
}
