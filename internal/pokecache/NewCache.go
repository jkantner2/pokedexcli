package pokecache

import (
	"time"
	"sync"
)

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		entryList:	make(map[string]cacheEntry),
		mu:	&sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}
