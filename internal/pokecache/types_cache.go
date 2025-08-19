package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	entryList	map[string]cacheEntry
	mu 		*sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()	
	c.entryList[key] = cacheEntry{
		createdAt:	time.Now().UTC(),
		val:		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.entryList[key]; ok {
		return c.entryList[key].val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.entryList {
		if value.createdAt.Before(now.Add(-interval)) {
			delete(c.entryList, key)
		}
	}
}
