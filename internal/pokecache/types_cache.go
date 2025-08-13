package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	entryList	map[string]cacheEntry
	mu 		sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.entryList == nil {
		c.entryList = make(map[string]cacheEntry)
	}
	c.entryList[key] = cacheEntry{}
	entry := c.entryList[key]
	entry.val = val
	c.entryList[key] = entry
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
