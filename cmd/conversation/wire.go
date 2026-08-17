//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"log/slog"

	"github.com/velony-app/conversation/internal/application"
	"github.com/velony-app/conversation/internal/conf"
	"github.com/velony-app/conversation/internal/infrastructure"
	"github.com/velony-app/conversation/internal/presentation"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Infrasturcture, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		infrastructure.ProviderSet,
		presentation.ProviderSet,
		application.ProviderSet,
		newApp,
	))
}
