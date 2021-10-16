package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (web *Web) ProjectPage(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return "", err
	}
	project, err := web.s.GetProjectByID(id)
	http.Redirect(w, r, project.RedirectURL, 302)
	return "", nil
}
