package main

import (
	"log/slog"
	"sync"
)

func main() {
	db := &DB{
		cache: NewCache(),
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(req int) {
			defer wg.Done()
			data := db.GetArticle(req, 1)
			slog.Info("data", "data", data, "req", req)
		}(i)
	}
	wg.Wait()
}

type DB struct {
	cache *Cache
}

func (db *DB) GetArticle(req int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	data = &Article{
		ID:      id,
		Content: "FooBar",
	}
	db.cache.Set(id, data)
	slog.Info("missing cache", "id", id, "req", req)
	return data
}
