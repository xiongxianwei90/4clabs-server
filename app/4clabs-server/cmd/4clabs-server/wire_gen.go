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
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/nft"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/trade"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/usecase"
	"4clabs-server/pkg/auth"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger, jwtUtils *auth.JwtUtils, contextUtils *auth.ContextUtils) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	nftNft := nft.NewNft(dataData)
	nftgoService := nftgo.NewService(bootstrap, nftNft, dataData)
	address := usecase.NewAddress(nftgoService)
	register := repo.NewRegister(nftgoService, dataData, nftNft)
	usecaseNft := usecase.NewNft(nftgoService, register)
	user := repo.NewUser(dataData)
	usecaseAuth := usecase.NewAuth(user, register, jwtUtils)
	ticket := repo.NewTicket(dataData)
	usecaseTicket := usecase.NewTicket(ticket)
	comic := repo.NewComic(dataData, nftgoService)
	tradeTrade := trade.NewTrade(dataData, nftgoService)
	comics := usecase.NewComics(comic, tradeTrade)
	script := repo.NewScript(bootstrap, dataData, nftgoService, register)
	usecaseScript := usecase.NewScript(script)
	serviceService := service.NewService(address, usecaseNft, usecaseAuth, contextUtils, usecaseTicket, comics, usecaseScript)
	httpServer := server.NewHTTPServer(bootstrap, serviceService, logger, contextUtils, jwtUtils)
	grpcServer := server.NewGRPCServer(bootstrap, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
