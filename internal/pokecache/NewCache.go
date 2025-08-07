package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) (Cache) {
	newCache := Cache{}
	newCache.reapLoop(interval)
	return newCache
}
