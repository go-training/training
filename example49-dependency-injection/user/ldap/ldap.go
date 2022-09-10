package ldap

import (
	"github.com/go-training/example49-dependency-injection/cache"
	"github.com/go-training/example49-dependency-injection/config"
)

type Service struct {
	bindUsername string
	bindPassword string
	cache        *cache.Service
}

func New(cfg config.Config, cache *cache.Service) (*Service, error) {
	return &Service{
		bindUsername: cfg.Ldap.BindUsername,
		bindPassword: cfg.Ldap.BindPassword,
		cache:        cache,
	}, nil
}
