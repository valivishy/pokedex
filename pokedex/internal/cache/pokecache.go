package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	Mutex    sync.RWMutex
	Interval time.Duration
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entries:  make(map[string]cacheEntry),
		Interval: interval,
	}

	go func() {
		cache.reapLoop()
	}()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()

	cache.Entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()

	entry, ok := cache.Entries[key]
	return entry.val, ok
}

func (cache *Cache) reapLoop() {
	timer := time.NewTicker(cache.Interval)
	for t := range timer.C {
		cache.reap(t)
	}
}

// region PRIVATE
func (cache *Cache) reap(tick time.Time) {
	cache.Mutex.RLock()
	defer cache.Mutex.RUnlock()

	for index, entry := range cache.Entries {
		if entry.createdAt.Add(cache.Interval).Before(tick) {
			fmt.Printf("%s expired, deleting\n", index)
			delete(cache.Entries, index)
		}
	}
}

//endregion PRIVATE
