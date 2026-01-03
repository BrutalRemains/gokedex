package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(1)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestCreateCache2(t *testing.T) {
	cache := NewCache(1)

	cache.Add("key1", []byte("val1"))
	actual, ok := cache.Get("key1")
	if !ok {
		t.Error("key1 not found")
	}
	if string(actual) != "val1" {
		t.Error("value doesn't match")
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been deleted", keyOne)
	}
}
