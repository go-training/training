package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	// Config provides the system configuration.
	Config struct {
		Logging    Logging
		Server     Server
		Prometheus Prometheus
		Ldap       Ldap
		Crowd      Crowd
		Cache      Cache
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug   bool   `envconfig:"APP_LOGS_DEBUG"`
		Level   string `envconfig:"APP_LOGS_LEVEL" `
		NoColor bool   `envconfig:"APP_LOGS_COLOR"`
		Pretty  bool   `envconfig:"APP_LOGS_PRETTY"`
		Text    bool   `envconfig:"APP_LOGS_TEXT"`
	}

	// Server provides the server configuration.
	Server struct {
		Addr  string `envconfig:"-"`
		Host  string `envconfig:"APP_SERVER_HOST" default:"localhost:8080"`
		Port  string `envconfig:"APP_SERVER_PORT" default:":8080"`
		Proto string `envconfig:"APP_SERVER_PROTO" `
		Pprof bool   `envconfig:"APP_PPROF_ENABLED"`
		Acme  bool   `envconfig:"APP_TLS_AUTOCERT"`
		Email string `envconfig:"APP_TLS_EMAIL"`
		Cert  string `envconfig:"APP_TLS_CERT"`
		Key   string `envconfig:"APP_TLS_KEY"`
		Debug bool   `envconfig:"APP_SERVER_DEBUG"`
	}

	// Prometheus logger
	Prometheus struct {
		Token string `envconfig:"APP_PROMETHEUS_TOKEN"`
	}

	Ldap struct {
		Host         string
		Port         int
		BindUsername string
		BindPassword string
		TestUsername string
		TestPassword string
	}

	// Crowd config
	Crowd struct {
		BasicUsername string `envconfig:"APP_CROWD_BASIC_USERNAME"`
		BasicPassword string `envconfig:"APP_CROWD_BASIC_PASSWORD"`
		TestUsername  string `envconfig:"APP_CROWD_TEST_USERNAME"`
		TestPassword  string `envconfig:"APP_CROWD_TEST_PASSWORD"`
	}

	Cache struct {
		DefaultExpiration time.Duration `envconfig:"APP_CACHE_DEFAULT_EXPIRATION"`
		CleanupInterval   time.Duration `envconfig:"APP_CACHE_CLEANUP_INTERVAL"`
	}
)

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}
