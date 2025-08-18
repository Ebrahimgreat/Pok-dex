package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]CacheEntry
	mutex    sync.Mutex
	interval time.Duration
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{}

	c.data = make(map[string]CacheEntry)
	c.interval = interval
	go c.reapLoop()
	return &c

}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > c.interval {

				delete(c.data, key)
			}
		}
		c.mutex.Unlock()

	}
}
func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry := CacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
	c.data[key] = entry
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.data[key]
	if ok {
		return value.val, true
	}
	return []byte{}, false

}
