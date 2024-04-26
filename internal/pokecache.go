package internal

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
	print("Creating new cache")
	cache := &Cache{
		cacheEntry: make(map[string]cacheEntry),
	}

	ticker := time.Ticker{
		C: make(chan time.Time),
	}

	_, ok := <-ticker.C
	println("Starting ticker, ok:", ok)
	if ok {
		println("Starting ticker")
		go cache.reapLoop(interval)
	}

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
			println("Reaping key", key)
			delete(cache.cacheEntry, key)
		}
	}
	cache.mutex.Unlock()

}
