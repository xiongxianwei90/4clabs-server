//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"4clabs-server/pkg/auth"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, *auth.JwtUtils, *auth.ContextUtils) (*kratos.App, func(), error) {
	panic(wire.Build(adapter.Provider, usecase.ProvderSet, newApp))
}
