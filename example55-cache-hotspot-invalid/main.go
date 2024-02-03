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
			_ = db.GetArticle(req, 1)
		}(i)
	}
	wg.Wait()
}

type DB struct {
	cache *Cache
}

func (db *DB) GetArticle(worker int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "worker", worker)
		return data
	}

	data = &Article{
		ID:      id,
		Content: "FooBar",
	}
	db.cache.Set(id, data)
	slog.Info("missing cache", "id", id, "worker", worker)
	return data
}
