package web

import (
	"net/http"
	"tochkru-golang/internal/app/templates"
)

func (web *Web) IndexPage(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return templates.IndexPage(), nil
}
