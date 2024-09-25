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
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{val: val, createdAt: time.Now()}

}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}
func (ch *Cache) ReapLoop() {
	ticker := time.NewTicker(ch.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		ch.mux.Lock()
		for key, entry := range ch.cache {
			if time.Since(entry.createdAt) > ch.interval {
				delete(ch.cache, key)
			}
		}
		ch.mux.Unlock()
	}
}
