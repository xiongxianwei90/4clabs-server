//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"4clabs-server/app/contractlistener/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger) (*ContractScriptApplication, func(), error) {
	panic(wire.Build(NewCronJobApplication))
}
