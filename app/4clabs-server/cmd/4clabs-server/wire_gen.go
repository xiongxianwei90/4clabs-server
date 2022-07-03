// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driven/server"
	"4clabs-server/app/4clabs-server/internal/adapter/driven/service"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/usecase"
	"4clabs-server/pkg/auth"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger, jwtUtils *auth.JwtUtils) (*kratos.App, func(), error) {
	nftgoService := nftgo.NewService(bootstrap)
	address := usecase.NewAddress(nftgoService)
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	register := repo.NewRegister(dataData)
	nft := usecase.NewNft(nftgoService, register)
	user := repo.NewUser(dataData)
	usecaseAuth := usecase.NewAuth(user, register, jwtUtils)
	ticket := repo.NewTicket(dataData)
	usecaseTicket := usecase.NewTicket(ticket)
	serviceService := service.NewService(address, nft, usecaseAuth, usecaseTicket)
	httpServer := server.NewHTTPServer(bootstrap, serviceService, logger)
	grpcServer := server.NewGRPCServer(bootstrap, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
