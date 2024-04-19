package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      *sync.Mutex
	entries map[string]cacheEntry

	ttl time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		mu:      &sync.Mutex{},
		entries: make(map[string]cacheEntry),

		ttl: ttl,
	}

	go cache.reapLoop()

	return cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}

	if time.Since(entry.createdAt) > cache.ttl {
		delete(cache.entries, key)
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.ttl)
	for range ticker.C {
		cache.reap()
	}
}

func (cache *Cache) reap() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	for key, entry := range cache.entries {
		if time.Since(entry.createdAt) > cache.ttl {
			delete(cache.entries, key)
		}
	}
}
