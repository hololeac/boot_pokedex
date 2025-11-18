package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var res []byte
	c.mu.Lock()
	entrie, ok := c.entries[key]
	if !ok {
		res = nil
	}
	c.mu.Unlock()
	res = entrie.val
	return res, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		timeToDie := time.Now().Add(-interval)
		for key := range c.entries {
			if c.entries[key].createdAt.Before(timeToDie) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
