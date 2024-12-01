package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
}

// NewCache -
type cacheEntry struct {
	val []byte
	createdAt time.Time
}
