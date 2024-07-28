package main

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

func (cache cache) add(key string, ce cacheEntry) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cache[key] = ce
	cache.reap()
}

func (cache cache) get(key string) cacheEntry {
	cache.reap()
	return cache.cache[key]
}

func (cache cache) reap() {
	for key, cacheItem := range cache.cache {
		if time.Since(cacheItem.createdAt) > 60*time.Second {
			delete(cache.cache, key)
		}
	}
}
