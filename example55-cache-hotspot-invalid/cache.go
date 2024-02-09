package main

import "sync"

type Cache[K comparable, V any] struct {
	sync.Mutex
	entries map[K]V
}

func (c *Cache[K, V]) Get(id K) (v V) {
	if _, ok := c.entries[id]; !ok {
		return v
	}
	c.Lock()
	defer c.Unlock()
	return c.entries[id]
}

func (c *Cache[K, V]) Set(id K, article V) {
	c.Lock()
	defer c.Unlock()
	c.entries[id] = article
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		entries: make(map[K]V),
	}
}
