package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{createdAt: time.Now().UTC(), val: val}
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entries, k)
		}
	}
}