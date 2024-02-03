package example55

import (
	"sync"
)

func main() {
	db := &DB{
		cache: NewCache(),
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			_ = db.GetArticle(1)
		}()
	}
	wg.Wait()
}

type DB struct {
	cache *Cache
}

func (db *DB) GetArticle(id int) *Article {
	data := db.cache.Get(id)
	if data != nil {
		return data
	}

	data = &Article{
		ID:      id,
		Content: "FooBar",
	}
	db.cache.Set(id, data)

	return data
}
