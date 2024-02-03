package main

import (
	"log/slog"
	"sync"

	"golang.org/x/sync/singleflight"
)

func main() {
	db := &DB{
		cache:  NewCache(),
		engine: singleflight.Group{},
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(req int) {
			defer wg.Done()
			data := db.GetArticle(req, 1)
			slog.Info("data info", "data", data, "req", req)
		}(i)
	}
	wg.Wait()
}

type DB struct {
	cache  *Cache
	engine singleflight.Group
}

func (db *DB) GetArticle(req int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	row, _, shared := db.engine.Do("article", func() (interface{}, error) {
		slog.Info("missing cache", "id", id, "req", req)
		data := &Article{
			ID:      id,
			Content: "FooBar",
		}
		db.cache.Set(id, data)
		return data, nil
	})

	slog.Any("shared", shared)

	return row.(*Article)
}
