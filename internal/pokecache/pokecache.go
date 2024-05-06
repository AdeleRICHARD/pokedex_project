package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mutex      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]cacheEntry),
	}

	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			cache.reapLoop(interval) // Appelez votre fonction de nettoyage
		}
	}()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	cache.cacheEntry[key] = cacheEntry{
		val: val,
	}
	cache.mutex.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	val := cache.cacheEntry[key].val
	if val != nil {
		return val, true
	}
	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	cache.mutex.Lock()
	for key, timeEntry := range cache.cacheEntry {
		if time.Since(timeEntry.createdAt) > interval {
			delete(cache.cacheEntry, key)
		}
	}
	cache.mutex.Unlock()

}
