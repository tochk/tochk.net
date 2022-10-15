package web

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"tochkru-golang/internal/app/metrics"
	"tochkru-golang/internal/app/service"
	"tochkru-golang/internal/app/templates"
)

type Web struct {
	s                    *service.Service
	m                    *metrics.Metrics
	currentRequestsCount int64
	cachePurge           sync.Mutex
	mu                   sync.RWMutex
	cache                map[string]string
	cacheEnabled         bool
}

const (
	LanguageRu = "ru"
	LanguageEn = "en"
)

func New(s *service.Service, m *metrics.Metrics, cache bool) *Web {
	return &Web{
		s:                    s,
		m:                    m,
		currentRequestsCount: 0,
		cachePurge:           sync.Mutex{},
		cacheEnabled:         cache,
		mu:                   sync.RWMutex{},
		cache:                map[string]string{},
	}
}

func (web *Web) Wrapper(f func(w http.ResponseWriter, r *http.Request) (string, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		web.cachePurge.Lock()
		//nolint:staticcheck
		web.cachePurge.Unlock()
		atomic.AddInt64(&web.currentRequestsCount, 1)
		defer atomic.AddInt64(&web.currentRequestsCount, -1)
		t := time.Now()
		log.Infof("Requested [%s] %s page from %s", r.Method, r.URL.Path, r.Header.Get("X-Real-IP"))
		defer func() {
			log.Debugf("Served [%s] %s page from %s for %v", r.Method, r.URL.Path, r.Header.Get("X-Real-IP"), time.Since(t))
		}()
		if !strings.HasPrefix(r.URL.Path, "/admin/") &&
			!strings.HasPrefix(r.URL.Path, "/project/") {
			page, found := web.GetCachedPage(r.URL.Path)
			if found {
				log.Debugln("cache hit")
				fmt.Fprint(w, page)
				return
			}
		}
		resp, err := f(w, r)
		if err != nil {
			log.Errorln(err)
			fmt.Fprint(w, templates.ErrorPage(err))
			return
		}
		if !strings.HasPrefix(r.URL.Path, "/admin/") &&
			!strings.HasPrefix(r.URL.Path, "/project/") {
			web.SaveCache(r.URL.Path, resp)
		}
		fmt.Fprint(w, resp)

	}
}
