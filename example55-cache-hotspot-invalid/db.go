package main

import (
	"fmt"
	"log/slog"
	"time"

	gsingleflight "example55/singleflight"

	"golang.org/x/sync/singleflight"
)

type DB struct {
	cache  *Cache[int, *Article]
	engine singleflight.Group
}

func (db *DB) GetArticle(req int, id int) *Article {
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

func NewDB() *DB {
	return &DB{
		cache:  NewCache[int, *Article](),
		engine: singleflight.Group{},
	}
}

type DBG struct {
	cache  *Cache[int, *Article]
	engine gsingleflight.Group[int, *Article]
}

func (db *DBG) GetArticle(req int, id int) *Article {
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

func (db *DBG) GetArticleDo(req int, id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	row, err, _ := db.engine.Do(id, func() (*Article, error) {
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

	return row
}

func (db *DBG) GetArticleDoChan(req int, id int, t time.Duration) *Article {
	data := db.cache.Get(id)
	if data != nil {
		slog.Info("cache hit", "id", id, "req", req)
		return data
	}

	dataChan := db.engine.DoChan(id, func() (*Article, error) {
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
		return res.Val
	}
}

func NewDBG() *DBG {
	return &DBG{
		cache:  NewCache[int, *Article](),
		engine: gsingleflight.Group[int, *Article]{},
	}
}
