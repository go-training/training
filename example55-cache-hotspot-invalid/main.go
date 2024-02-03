package main

import (
	"fmt"
	"log/slog"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	db := &DB{
		cache:  NewCache(),
		engine: singleflight.Group{},
	}

	slog.Info("================== without singleflight ===================")

	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(req int) {
			defer wg.Done()
			data := db.GetArticleOld(req, 1)
			slog.Info("data info", "data", data, "req", req)
		}(i)
	}
	wg.Wait()

	slog.Info("================== using singleflight ===================")

	wg.Add(5)
	for i := 5; i < 10; i++ {
		go func(req int) {
			defer wg.Done()
			data := db.GetArticleNew(req, 2)
			slog.Info("data info", "data", data, "req", req)
		}(i)
	}
	wg.Wait()
}

type DB struct {
	cache  *Cache
	engine singleflight.Group
}

func (db *DB) GetArticleOld(req int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	slog.Info("missing cache", "id", id, "req", req)
	data = &Article{
		ID:      id,
		Content: "FooBar",
	}
	db.cache.Set(id, data)

	return data
}

func (db *DB) GetArticleNew(req int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	key := fmt.Sprintf("article:%d", id)
	row, err, _ := db.engine.Do(key, func() (interface{}, error) {
		slog.Info("missing cache", "id", id, "req", req)
		data := &Article{
			ID:      id,
			Content: "FooBar",
		}
		db.cache.Set(id, data)
		time.Sleep(100 * time.Millisecond)
		return data, nil
	})

	if err != nil {
		slog.Error("singleflight error", "err", err)
		return nil
	}

	return row.(*Article)
}
