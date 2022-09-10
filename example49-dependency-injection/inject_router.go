package main

import (
	"net/http"

	"github.com/go-training/example49-dependency-injection/config"
	"github.com/go-training/example49-dependency-injection/router"
	"github.com/go-training/example49-dependency-injection/user"

	"github.com/google/wire"
)

var routerSet = wire.NewSet( //nolint:deadcode,unused,varcheck
	provideRouter,
)

func provideRouter(
	cfg config.Config,
	user *user.Service,
) http.Handler {
	return router.New(cfg, user)
}
