package web

import (
	log "github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

func (web *Web) GetCachedPage(path string) (string, bool) {
	if !web.cacheEnabled {
		return "", false
	}
	web.mu.RLock()
	defer web.mu.RUnlock()
	cachedPage, found := web.cache[path]
	return cachedPage, found
}

func (web *Web) SaveCache(path, body string) {
	if !web.cacheEnabled {
		return
	}
	web.mu.Lock()
	defer web.mu.Unlock()
	web.cache[path] = body
}

func (web *Web) PurgeCache() {
	web.cachePurge.Lock()
	defer web.cachePurge.Unlock()
	for {
		if atomic.LoadInt64(&web.currentRequestsCount) == 0 {
			break
		}
		log.Debug("waiting for all requests to be done")
		time.Sleep(time.Millisecond * 100)
	}
	web.mu.Lock()
	defer web.mu.Unlock()

	web.cache = make(map[string]string, 0)

	log.Debug("cache purge successful")
}