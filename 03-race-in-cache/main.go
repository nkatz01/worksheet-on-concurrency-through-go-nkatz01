package main

import (
	"container/list"
	"sync"

)

// CacheSize determines how big the cache can grow
const CacheSize = 100

// KeyStoreCacheLoader is an interface for the KeyStoreCache
type KeyStoreCacheLoader interface {
	// Load implements a function for obtaining the data for the cache.
	Load(string) string
}

// KeyStoreCache is a LRU (Least-Recently-Used) cache of string key-value pairs.
type KeyStoreCache struct {
	mu sync.Mutex
	cache map[string]string
	pages list.List
	load  func(string) string
}

// New creates a new KeyStoreCache.
func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{

		load:  load.Load,
		cache: make(map[string]string),
	}
}

// Get gets the key from cache, loading it from the source if needed.
func (k *KeyStoreCache) Get(key string) string {
	k.mu.Lock()
	val, ok := k.cache[key]

	// Miss - load from database and save it in cache.
	if !ok {
		val = k.load(key)

		k.cache[key] = val

		k.pages.PushFront(key)

		// if cache is full remove the least used item.

		if len(k.cache) > CacheSize {
			delete(k.cache, k.pages.Back().Value.(string))
			k.pages.Remove(k.pages.Back())
		}


	}
	k.mu.Unlock()
	return val
}

// Loader implements KeyStoreLoader.
type Loader struct {
	DB *MockDB
}

// Load gets the data from the database.
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func run() *KeyStoreCache {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache)

	return cache
}

func main() {
	run()
}
