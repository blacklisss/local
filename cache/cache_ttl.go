package cache

import (
	"sync"
	"time"
)

// CacheItem represents an item in the cache.
type CacheItem struct {
	data       interface{}
	expiration time.Time
}

// Cache represents an in-memory cache with TTL.
type Cache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
}

// NewCache creates a new Cache instance.
func NewCache() *Cache {
	cache := &Cache{
		items: make(map[string]CacheItem),
	}
	go cache.startCleanupRoutine()
	return cache
}

// Set adds or updates an item in the cache with a TTL.
func (c *Cache) Set(key string, data interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	expiration := time.Now().Add(ttl)
	c.items[key] = CacheItem{
		data:       data,
		expiration: expiration,
	}
}

// Get retrieves an item from the cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	item, exists := c.items[key]
	c.mutex.RUnlock()

	if !exists {
		return nil, false
	}

	if time.Now().After(item.expiration) {
		// Item has expired, remove it from the cache
		c.mutex.Lock()
		delete(c.items, key)
		c.mutex.Unlock()
		return nil, false
	}
	return item.data, true
}

// startCleanupRoutine periodically removes expired items from the cache.
func (c *Cache) startCleanupRoutine() {
	cleanupInterval := 5 * time.Minute // Adjust as needed
	ticker := time.NewTicker(cleanupInterval)
	for range ticker.C {
		c.mutex.Lock()
		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}
