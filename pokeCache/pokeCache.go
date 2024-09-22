package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mux      *sync.RWMutex
	interval time.Duration
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		cache:    map[string]cacheEntry{},
		mux:      &sync.RWMutex{},
		interval: interval,
	}
}
func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{val: val, createdAt: time.Now()}

}
