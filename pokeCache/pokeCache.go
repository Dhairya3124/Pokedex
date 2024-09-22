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