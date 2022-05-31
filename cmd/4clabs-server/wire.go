// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"4clabs-server/internal/adapter/driven/server"
	"4clabs-server/internal/biz"
	"4clabs-server/internal/conf"
	"4clabs-server/internal/data"
	"4clabs-server/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
