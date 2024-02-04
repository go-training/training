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

	slog.Info("================== using singleflight Do ===================")

	wg.Add(5)
	for i := 5; i < 10; i++ {
		go func(req int) {
			defer wg.Done()
			data := db.GetArticleDo(req, 2)
			slog.Info("data info", "data", data, "req", req)
		}(i)
	}
	wg.Wait()

	slog.Info("================== using singleflight DoChan ===================")

	wg.Add(5)
	for i := 10; i < 15; i++ {
		go func(req int) {
			defer wg.Done()
			t := time.Duration(time.Duration(int64(req*10)) * time.Millisecond)
			data := db.GetArticleDoChan(req, 3, t)
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
	time.Sleep(100 * time.Millisecond)

	return data
}

func (db *DB) GetArticleDo(req int, id int) *Article {
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

func (db *DB) GetArticleDoChan(req int, id int, t time.Duration) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	key := fmt.Sprintf("article:%d", id)
	dataChan := db.engine.DoChan(key, func() (interface{}, error) {
		slog.Info("missing cache", "id", id, "req", req)
		data := &Article{
			ID:      id,
			Content: "FooBar",
		}
		db.cache.Set(id, data)
		time.Sleep(115 * time.Millisecond)
		return data, nil
	})

	select {
	case <-time.After(t):
		slog.Info("timeout", "id", id, "req", req)
		return nil
	case res := <-dataChan:
		return res.Val.(*Article)
	}
}
