package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GetReadyResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func (a *Api) GetReadyHandler(w http.ResponseWriter, r *http.Request) {
	if err := a.s.Ready(); err != nil {
		w.WriteHeader(500)
		log.Error("error on /ready: ", err)
		data, err := json.Marshal(GetReadyResponse{
			Status:  "error",
			Service: "tochknet",
		})
		if err != nil {
			log.Error("can't unmarshal json: ", err)
		}
		_, _ = w.Write(data)
		return
	}

	data, err := json.Marshal(GetReadyResponse{
		Status:  "ok",
		Service: "tochknet",
	})
	if err != nil {
		log.Error("can't unmarshal json: ", err)
	}

	_, _ = w.Write(data)
}
