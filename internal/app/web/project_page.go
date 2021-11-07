package web

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (web *Web) ProjectPageHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectPageHandler")
	}
	project, err := web.s.GetProjectByID(id)
	if err != nil {
		return "", errors.Wrap(err, "web.ProjectPageHandler")
	}
	http.Redirect(w, r, project.RedirectURL, 302)
	return "", nil
}
