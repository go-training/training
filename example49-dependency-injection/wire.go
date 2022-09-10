//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-training/example49-dependency-injection/config"

	"github.com/google/wire"
)

func InitializeApplication(cfg config.Config) (*application, error) {
	wire.Build(
		routerSet,
		userSet,
		newApplication,
	)
	return &application{}, nil
}
