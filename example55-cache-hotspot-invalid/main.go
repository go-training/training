package main

import (
	"log/slog"
	"sync"
	"time"
)

func SingleFlight(db Middleware) {
	slog.Info("================== without singleflight ===================")

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

func main() {
	slog.Info("singleflight (none Generic) test started..")
	SingleFlight(NewDB())
	slog.Info("singleflight (none Generic) test finished..")
	slog.Info("singleflight (Generic) test started..")
	SingleFlight(NewDBG())
	slog.Info("singleflight (Generic) test finished..")
}
