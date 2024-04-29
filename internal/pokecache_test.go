package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	t.Run("Test add to cache", func(t *testing.T) {
		cache := NewCache(interval)
		cache.Add("https://example.com", []byte("testdata"))
	})

	t.Run("Test get from cache", func(t *testing.T) {
		cache := NewCache(interval)
		cache.Add("https://example.com", []byte("testdata"))

		val, ok := cache.Get("https://example.com")
		assert.True(t, ok, "expected key to be present")
		assert.Equal(t, []byte("testdata"), val)
	})
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	t.Run("ReapLoop delete key", func(t *testing.T) {
		cache := NewCache(baseTime)
		cache.Add("https://example.com", []byte("testdata"))

		_, ok := cache.Get("https://example.com")
		assert.True(t, ok, "expected key to be present")

		time.Sleep(waitTime)

		_, ok = cache.Get("https://example.com")
		assert.True(t, !ok, "expected key to be deleted")
	})
}
