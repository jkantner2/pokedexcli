package pokecache

import (
	"sync"
	"time"

	"golang.org/x/tools/go/analysis/passes/defers"
)

type Cache struct{
	entryList	map[string]*cacheEntry
	mu 		sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entryList[key] = &cacheEntry{}
	c.entryList[key].val = val
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.entryList[key]; ok {
		return c.entryList[key].val, true
	} else {
		return nil, false
	}
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Reset(interval)
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.entryList {
		if time.Since(value.createdAt) >= interval {
			delete(c.entryList, key)
		} else {
			continue
		}
	}

}
