package web

func (web *Web) GetCachedPage(path string) (string, bool) {
	web.mu.RLock()
	defer web.mu.RUnlock()
	cachedPage, found := web.cache[path]
	return cachedPage, found
}

func (web *Web) SaveCache(path, body string) {
	web.mu.Lock()
	defer web.mu.Unlock()
	web.cache[path] = body
}
