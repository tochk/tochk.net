package web

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"tochkru-golang/internal/app/metrics"
	"tochkru-golang/internal/app/service"
	"tochkru-golang/internal/app/templates"
)

type Web struct {
	s     *service.Service
	m     *metrics.Metrics
	cache sync.Map
}

func New(s *service.Service, m *metrics.Metrics) *Web {
	return &Web{
		s:     s,
		m:     m,
		cache: sync.Map{},
	}
}

func (web *Web) Wrapper(f func(w http.ResponseWriter, r *http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Loaded %s page from %s", r.URL.Path, r.Header.Get("X-Real-IP"))
		resp, err := f(w, r)
		if err != nil {
			log.Errorln(err)
			fmt.Fprint(w, templates.ErrorPage(err))
			return
		}
		if resp != nil {
			fmt.Fprint(w, resp)
			return
		}
	}
}
