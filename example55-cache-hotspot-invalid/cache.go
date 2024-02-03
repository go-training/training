package main

import "sync"

type Cache struct {
	sync.Mutex
	entries map[int]*Article
}

func (c *Cache) Get(id int) *Article {
	if _, ok := c.entries[id]; !ok {
		return nil
	}
	c.Lock()
	defer c.Unlock()
	return c.entries[id]
}

func (c *Cache) Set(id int, article *Article) {
	c.Lock()
	defer c.Unlock()
	c.entries[id] = article
}

func NewCache() *Cache {
	return &Cache{
		entries: make(map[int]*Article),
	}
}
