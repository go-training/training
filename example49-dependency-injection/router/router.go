package router

import (
	"net/http"
	"path"

	"github.com/go-training/example49-dependency-injection/api"
	"github.com/go-training/example49-dependency-injection/config"
	"github.com/go-training/example49-dependency-injection/router/prometheus"
	"github.com/go-training/example49-dependency-injection/user"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func New(
	cfg config.Config,
	u *user.Service,
	middleware ...gin.HandlerFunc,
) http.Handler {
	if cfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.MaxMultipartMemory = 1024 << 20 // 1024 MiB

	r.Use(gin.Recovery())

	r.Use(logger.SetLogger(
		logger.WithUTC(true),
	))
	r.Use(middleware...)

	if cfg.Server.Pprof {
		pprof.Register(
			r,
			path.Join("debug", "pprof"),
		)
	}

	// 404 not found
	r.NoRoute(api.NotFound)

	r.GET("/metrics", prometheus.Handler(cfg.Prometheus.Token))
	r.GET("/healthz", api.Heartbeat)

	return r
}
