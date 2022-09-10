package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/go-training/example49-dependency-injection/config"
	"github.com/go-training/example49-dependency-injection/user"

	"github.com/appleboy/graceful"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

type application struct {
	router http.Handler
	user   *user.Service
}

func newApplication(
	router http.Handler,
	user *user.Service,
) *application {
	return &application{
		router: router,
		user:   user,
	}
}

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()

	_ = godotenv.Load(envfile)
	cfg, err := config.Environ()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("invalid configuration")
	}

	app, err := InitializeApplication(cfg)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("invalid configuration")
	}

	if ok := app.user.Login("test", "test"); !ok {
		log.Fatal().
			Err(err).
			Msg("invalid configuration")
	}

	m := graceful.NewManager()
	srv := &http.Server{
		Addr:              cfg.Server.Port,
		Handler:           app.router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	m.AddRunningJob(func(ctx context.Context) error {
		log.Info().Msgf("api server running on %s port", cfg.Server.Port)
		return listenAndServe(srv, cfg.Server.Cert, cfg.Server.Key)
	})

	m.AddShutdownJob(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)
	})

	<-m.Done()
}

func listenAndServe(s *http.Server, certPath string, keyPath string) error {
	if certPath != "" && keyPath != "" {
		return s.ListenAndServeTLS(certPath, keyPath)
	}

	return s.ListenAndServe()
}
