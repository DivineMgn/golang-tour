package main

import "sync"

// URLCache - url cache
type URLCache struct {
	data map[string]bool
	lock sync.Mutex
}

// Contains - check url
func (cache *URLCache) Contains(url string) (exists bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	_, exists = cache.data[url]
	return
}

// Add - append url to cache
func (cache *URLCache) Add(url string) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	cache.data[url] = true
}
