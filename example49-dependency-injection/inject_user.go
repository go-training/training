package main

import (
	"github.com/go-training/example49-dependency-injection/cache"
	"github.com/go-training/example49-dependency-injection/config"
	"github.com/go-training/example49-dependency-injection/user"
	"github.com/go-training/example49-dependency-injection/user/crowd"
	"github.com/go-training/example49-dependency-injection/user/ldap"

	"github.com/google/wire"
)

var userSet = wire.NewSet( //nolint:deadcode,unused,varcheck
	provideUser,
	provideLDAP,
	provideCROWD,
	provideCache,
)

func provideUser(
	l *ldap.Service,
	c *crowd.Service,
	cache *cache.Service,
) (*user.Service, error) {
	return user.New(l, c, cache)
}

func provideLDAP(
	cfg config.Config,
	cache *cache.Service,
) (*ldap.Service, error) {
	return ldap.New(cfg, cache)
}

func provideCROWD(
	cfg config.Config,
	cache *cache.Service,
) (*crowd.Service, error) {
	return crowd.New(cfg, cache)
}

func provideCache(
	cfg config.Config,
) (*cache.Service, error) {
	return cache.New(cfg)
}
