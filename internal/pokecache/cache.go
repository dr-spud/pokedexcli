// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cacheEntry, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.cacheMap {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.cacheMap, key)
		}
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.Clear()
	}
}
