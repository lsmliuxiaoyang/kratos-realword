//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-realwd/internal/biz"
	"kratos-realwd/internal/conf"
	"kratos-realwd/internal/data"
	"kratos-realwd/internal/server"
	"kratos-realwd/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.ProviderSet, biz.ProviderSet, service.ServiceProviderSet, newApp))
}
