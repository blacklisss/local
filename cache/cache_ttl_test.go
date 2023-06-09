package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	// Set an item with a TTL of 2 seconds
	cache.Set("key1", "value1", 2*time.Second)

	// Retrieve the item immediately after setting it
	value, exists := cache.Get("key1")
	if !exists {
		t.Error("Expected key1 to exist in the cache")
	}
	if value != "value1" {
		t.Errorf("Expected value1, but got %v", value)
	}

	// Wait for the TTL to expire
	time.Sleep(3 * time.Second)

	// Retrieve the item after the TTL has expired
	value, exists = cache.Get("key1")
	if exists {
		t.Error("Expected key1 to be expired and not exist in the cache")
	}
	if value != nil {
		t.Errorf("Expected nil value, but got %v", value)
	}

	// Set an item with a TTL of 5 seconds
	cache.Set("key2", "value2", 5*time.Second)

	// Wait for 3 seconds
	time.Sleep(3 * time.Second)

	// Retrieve the item before the TTL has expired
	value, exists = cache.Get("key2")
	if !exists {
		t.Error("Expected key2 to exist in the cache")
	}
	if value != "value2" {
		t.Errorf("Expected value2, but got %v", value)
	}

}
