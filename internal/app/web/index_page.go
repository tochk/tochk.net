package web

import (
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) IndexPage(w http.ResponseWriter, r *http.Request) (string, error) {

	return templates.IndexPage(), nil
}
