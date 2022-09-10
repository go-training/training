package crowd

import (
	"github.com/go-training/example49-dependency-injection/cache"
	"github.com/go-training/example49-dependency-injection/config"
)

type Service struct {
	basicUsername string
	basicPassword string
	cache         *cache.Service
}

func New(cfg config.Config, cache *cache.Service) (*Service, error) {
	return &Service{
		basicUsername: cfg.Crowd.BasicUsername,
		basicPassword: cfg.Crowd.BasicPassword,
		cache:         cache,
	}, nil
}
